package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"line_china/search_server/src/common/route"
	"line_china/search_server/src/provider"
	"line_china/search_server/src/service"
	"line_china/search_server/src/utils"
	"log"
)

func main() {
	var err error

	// 加载配置文件
	if err = utils.ConfigInit(); err != nil {
		log.Println(err.Error())
	}

	// 开启Nacos监听
	if err = utils.ListenConfig(); err != nil {
		fmt.Println(err)
	}

	// 初始化Redis
	//utils.RedisInit()

	// 初始化异步Kafka
	provider.InitAsyncKafka()
	// 计算域权重
	service.CalculateWeight()
	r := gin.Default()
	r = route.PathRoute(r) // 路由
	err = r.Run(":8082")
	if err != nil {
		return
	}

}
