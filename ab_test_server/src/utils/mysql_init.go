package utils

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var GormDb *gorm.DB

// MysqlInit 初始化gorm连接MySQL
func MysqlInit() {
	var err error
	log.Print("开始初始化MySQL")
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Data.Username,
		Config.Data.Password,
		Config.Data.Ip,
		Config.Data.Port,
		Config.Data.DataBase,
	)

	GormDb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       url,   // 账号密码地址端口
		DefaultStringSize:         100,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// ...gorm配置
	})
	if err != nil {
		log.Println(err)
		return
	}

	var sqlDB *sql.DB
	sqlDB, err = GormDb.DB() // 直接通过上面连接得到的db，调用DB()函数返回即可设置
	if err != nil {
		log.Println(err)
		return
	}

	sqlDB.SetMaxIdleConns(Config.Data.MaxIdleConns) // SetMaxIdleCons 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(Config.Data.MaxOpenConns) // SetMaxOpenCons 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour)             // SetConnMaxLifetime 设置了连接可复用的最大时间。

	//尝试连接数据库
	err = sqlDB.Ping()
	if err != nil {
		log.Println(err)
		return
	}

	log.Print("成功初始化MySQL")
}
