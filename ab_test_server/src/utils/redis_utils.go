package utils

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"time"
)

var RedisPool *redis.Pool
var L *RedisRWMutex

// RedisInit Redis初始化
func RedisInit() {
	log.Println("开始初始化Redis")
	RedisPool = &redis.Pool{
		MaxIdle:     Config.Redis.MaxIdle,     //最大空闲连接数
		MaxActive:   Config.Redis.MaxActive,   //在给定时间内，允许分配的最大连接数（当为零时，没有限制）
		IdleTimeout: Config.Redis.IdleTimeout, //在给定时间内，保持空闲状态的时间，若到达时间限制则关闭连接（当为零时，没有限制）
		//提供创建和配置应用程序连接的一个函数
		Dial: func() (redis.Conn, error) {
			log.Println("开始连接Redis")
			c, err := redis.Dial(Config.Redis.Network, Config.Redis.Ip+":"+strconv.Itoa(Config.Redis.Port))
			if err != nil {
				log.Println(err)
				return nil, err
			}
			//如果redis设置了用户密码，使用auth指令
			//if _, err = c.Do("AUTH", ""); err != nil {
			//	c.Close()
			//	return nil, err
			//}
			log.Println("成功连接Redis")
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	//L = NewRedisRWMutex("experiment_data_lock", 15, 10, 10*time.Millisecond)
	log.Println("成功初始化Redis")
}

// RedisRWMutex Redis读写锁基本配置结构体
type RedisRWMutex struct {
	key           string
	maxLockSecond int           // 锁定状态标记的最大时间
	maxRetryTimes int           // 最大阻塞重试次数
	retryInterval time.Duration // 重试的间隔
}

func NewRedisRWMutex(key string, maxLockSecond int, maxRetryTimes int, retryInterval time.Duration) *RedisRWMutex {
	return &RedisRWMutex{
		key:           key,
		maxLockSecond: maxLockSecond,
		maxRetryTimes: maxRetryTimes,
		retryInterval: retryInterval,
	}
}

// RLock 获取读锁
func (m *RedisRWMutex) RLock(conn redis.Conn) (int, error) {
	maxRetry := m.maxRetryTimes
	rs, e := m.rLockLoop(&maxRetry, conn)
	if e != nil {
		return 0, errorx.Wrap(e)
	}

	return rs, nil
}

// 循环获取读锁
func (m *RedisRWMutex) rLockLoop(retryTimes *int, conn redis.Conn) (int, error) {
	// 写锁定时, 锁状态置为2, 阻塞其他读写
	// 读锁定时,锁状态置为1,不做任何阻塞
	// 无锁时，锁状态为0，或者不存在该key
	// 返回3表示需要阻塞
	// 返回2表示可执行
	var script = `
        local stat = redis.call('GET',KEYS[1]);
        
        -- 不存在,无锁时,返回可执行，并标记为读锁中
        if not stat then
            redis.call('SETEX', KEYS[1],ARGV[1],1)
            return 2;
        end
        
        -- 存在，但是出于无锁状态,返回可执行，标记为读锁中
        if tonumber(stat) == 0 then
            redis.call('SETEX', KEYS[1],ARGV[1],1)
            return 2;
        end

        -- 写锁定时，返回阻塞
         if tonumber(stat) == 2 then
            return 3;
         end

         -- 读锁定时，返回放行
         if tonumber(stat) == 1 then
            return 2;
         end
 
         -- 预期之外的结果
         return 4;
`
	vint, e := redis.Int(conn.Do("eval", script, 1, m.key, m.maxLockSecond))
	if e != nil {
		return 0, errorx.Wrap(e)
	}

	// 可执行
	if vint == 2 {
		return 0, nil
	}

	if vint == 3 {
		*retryTimes--
		if *retryTimes == 0 {
			return 1, nil
		}

		time.Sleep(m.retryInterval)
		return m.rLockLoop(retryTimes, conn)
	}

	return 0, errorx.NewFromStringf("unexpected lock stat return %d", vint)
}

// RUnLock 释放读锁
func (m *RedisRWMutex) RUnLock(conn redis.Conn) error {
	_, err := conn.Do("del", m.key)
	return err
}

// Lock 0 可执行, 1 超出了最大重试次数了
func (m RedisRWMutex) Lock(conn redis.Conn) (int, error) {
	var max = m.maxRetryTimes

	rs, e := m.loopLock(&max, conn)
	if e != nil {
		return 0, errorx.Wrap(e)
	}

	return rs, nil
}

// 循环获取写锁
func (m *RedisRWMutex) loopLock(retryTimes *int, conn redis.Conn) (int, error) {

	// 写锁定时, 锁状态置为2, 阻塞其他读写
	// 读锁定时, 锁状态置为1, 不做任何阻塞
	// 无锁时, 锁状态为0, 或者不存在该key
	// 返回3表示需要阻塞
	// 返回2表示可执行
	var script = `
        local stat = redis.call('GET',KEYS[1]);
        
        -- 无锁时,返回可执行，并标记为写锁中
        if not stat then
            redis.call('SETEX', KEYS[1],ARGV[1],2)
            return 2;
        end
        
        -- 无锁,返回可执行，标记为写锁中
        if math.abs(tonumber(stat)) < 0.1 then
            redis.call('SETEX', KEYS[1],ARGV[1],2)
            return 2;
        end

        -- 写锁定时，返回阻塞
         if math.abs(tonumber(stat)-2) < 0.1 then
            return 3;
         end

         -- 读锁定时,返回阻塞
          if math.abs(tonumber(stat)-1) < 0.1 then
            return 3;
         end

         -- 预期之外的结果
         return 4;
`
	vint, e := redis.Int(conn.Do("eval", script, 1, m.key, m.maxLockSecond))
	if e != nil {
		return 0, errorx.Wrap(e)
	}

	// 可执行
	if vint == 2 {
		return 0, nil
	}

	if vint == 3 {
		*retryTimes--
		if *retryTimes == 0 {
			return 1, nil
		}
		time.Sleep(m.retryInterval)
		return m.loopLock(retryTimes, conn)
	}
	return 0, errorx.NewFromStringf("unexpected lock stat return %d", vint)
}

// UnLock 释放写锁
func (m *RedisRWMutex) UnLock(conn redis.Conn) error {
	_, err := conn.Do("del", m.key)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
