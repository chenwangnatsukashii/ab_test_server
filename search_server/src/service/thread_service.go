package service

import (
	"encoding/json"
	"fmt"
	comModel "line_china/common/model"
	"line_china/search_server/src/utils"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var Quit chan int // 退出通道，接收
var In chan int
var publish comModel.Publish

func worker(in, quit <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			fmt.Println("收到退出信号")
			return
		case <-in:
		}
	}
}

// StartUp 开始请求
func StartUp(number int) {
	Quit = make(chan int)
	In = make(chan int)
	wg.Add(1)

	publish = CalculateWeight()
	go worker(In, Quit)
	worker := utils.NewWorker(1, 1)

	for i := 0; i < number; i++ {
		userId := worker.GetID()
		In <- userId

		time.Sleep(1 * time.Second)

		// 流量的实验选择加上读锁
		utils.RwLock.RLock()
		ChooseExp(userId)
		utils.RwLock.RUnlock()
	}

}

// ShutDown 停止请求
func ShutDown(id int) {
	close(Quit)
	wg.Wait()
}

// IsImpression 是否停留
func IsImpression() bool {
	return rand.Intn(100) < 65
}

// IsClick 是否点击
func IsClick() bool {
	return rand.Intn(100) < 50
}

// IsOrder 是否下单
func IsOrder() bool {
	return rand.Intn(100) < 40
}

// ChooseExp 流量的实验选择
func ChooseExp(userId int) {
	code := userId % 100

	var domainArray []int // 获取域权重的累加和
	for i, domain := range publish.Domains {
		if i == 0 {
			domainArray = append(domainArray, domain.Weight)
		} else {
			domainArray = append(domainArray, domainArray[i-1]+domain.Weight)
		}
	}

	var resModel = comModel.ResultModel{} // 创建返回结果的实体
	resModel.PublishId = publish.PublishId

	for i := 0; i < len(domainArray); i++ {
		if code <= domainArray[i] { // 权重筛选流入的域
			layers := publish.Domains[i].Layers
			for _, layer := range layers {
				exps := layer.Exps
				var expArray []int // 获取实验权重的累加和
				for i, exp := range exps {
					if i == 0 {
						expArray = append(expArray, exp.Weight)
					} else {
						expArray = append(expArray, expArray[i-1]+exp.Weight)
					}
				}

				for i := 0; i < len(expArray); i++ { // 权重筛选流入的实验
					if code <= expArray[i] {
						resModel.ExpId = exps[i].Id
						resModel.UserId = userId
						resModel.IsImpression = IsImpression()
						if resModel.IsImpression {
							resModel.IsClick = IsClick()
							if resModel.IsClick {
								resModel.IsOrder = IsOrder()
							}
						}

						data, err := json.Marshal(resModel)
						if err != nil {
							log.Println(err)
						}

						// 向Kafka发送次流量在某一层选择的实验信息
						if err = SendMsgKafka(publish.PublishId, data); err != nil {
							log.Println(err)
						}
						break
					}
				}
			}
		}
	}

}
