package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"line_china/ab_test_server/src/common/router"
	"line_china/ab_test_server/src/consumer"
	"line_china/ab_test_server/src/service"
	"line_china/ab_test_server/src/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var err error

	utils.LogInit()                 // 初始化log打印地址
	utils.ConfigInit()              // 加载配置文件
	utils.MysqlInit()               // MySQL初始化
	consumer.NewKafka().KafkaInit() // KafkaGroup初始化，接受searchServer处理后的结果
	utils.InitNacos()               // Nacos初始化，用作实验发布的配置中心
	utils.RedisInit()               // 初始化Redis。做接口的幂等校验
	go service.StartTimer(5)        // 启动定时任务，定时保存内存中的统计结果

	r := gin.Default()
	r.Use(cors())           // Cors 解决前后端分离跨域问题
	r = router.PathRoute(r) // 添加路由规则

	//启动服务(优雅关机)
	srv := &http.Server{
		Addr:    utils.Config.Server.Port,
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) // 此处不会阻塞

	defer func() {
		if consumer.KafkaClient != nil && consumer.KafkaClient.Close() == nil {
			log.Println("Kafka关闭成功")
		} else {
			log.Println("Kafka关闭失败")
		}

		if utils.RedisPool != nil && utils.RedisPool.Close() == nil {
			log.Println("Redis关闭成功")
		} else {
			log.Println("Redis关闭失败")
		}
	}()

	for {
		select {
		case s := <-quit:
			log.Printf("recv exit signal: %s\n", s.String())
			// 创建一个5秒超时的context
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
			if err = srv.Shutdown(ctx); err != nil {
				log.Println("Server Shutdown err: ", err)
			}
			log.Println("Server Shutdown Success")
			return
		}
	}

}

func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
