package service

import (
	"line_china/ab_test_server/src/dao"
	"line_china/ab_test_server/src/model"
	comModel "line_china/common/model"
	"sync"
	"time"
)

const (
	PublishID    = "publish_id"
	ExpId        = "exp_id"
	IsImpression = "is_impression"
	IsClick      = "is_click"
	IsOrder      = "is_order"
	Total        = "total"
)

var expMap = struct {
	sync.RWMutex
	m map[int]map[string]int
}{m: make(map[int]map[string]int)}

// AddOneRecord 添加一条实验数据记录
func AddOneRecord(input comModel.ResultModel) {
	key := input.ExpId

	// 处理实验记录获取写锁
	expMap.Lock()
	_, ok := expMap.m[key]
	if !ok {
		countMap := map[string]int{IsImpression: 0, IsClick: 0, IsOrder: 0, Total: 0, PublishID: input.PublishId, ExpId: input.ExpId}
		expMap.m[key] = countMap
	}

	if input.IsImpression {
		expMap.m[key][IsImpression] += 1
	}
	if input.IsClick {
		expMap.m[key][IsClick] += 1
	}
	if input.IsOrder {
		expMap.m[key][IsOrder] += 1
	}
	expMap.m[key][Total] += 1

	// 处理实验记录释放写锁
	expMap.Unlock()

}

// FinalResult 获取当前实验的统计指标
func FinalResult() error {
	var err error

	// 保存实验记录获取读锁
	expMap.RLock()
	for _, m := range expMap.m {
		err = AddReport(m)
		if err != nil {
			break
		}
	}

	// 保存实验记录释放读锁
	expMap.RUnlock()
	return err
}

// AddReport 添加某一时刻的全量数据
func AddReport(m map[string]int) error {
	report := model.Report{}
	report.PublishId = m[PublishID]
	report.ExpId = m[ExpId]

	report.ImpressionCnt = m[IsImpression]
	report.ClickCnt = m[IsClick]
	report.OrderCnt = m[IsOrder]
	report.Total = m[Total]

	report.DateInfo = time.Now()
	return dao.AddReport(report)
}
