package service

import (
	"line_china/ab_test_server/src/dao"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
	"strconv"
)

// GetIndicatorById 添加
func GetIndicatorById(id int) ([]model.EffectRes, error) {
	var realRes []model.EffectRes

	expIndicator := model.Indicator{}
	controlIndicator := model.Indicator{}

	expEntity, err := SelectExp(id)
	if err != nil {
		return nil, err
	}

	controlEntity, err := dao.GetControl(expEntity.Pid)
	if err != nil {
		return nil, err
	}

	// 基本信息
	realRes = append(realRes, model.EffectRes{Param0: "小组", Param1: expEntity.Name, Param2: controlEntity.Name, Param3: "-", Param4: "-", Param5: "-"})

	// 获取实验统计结果的读锁
	expMap.RLock()
	expIndicator.ImpressionCnt, expIndicator.ClickCnt, expIndicator.OrderCnt, expIndicator.Total =
		expMap.m[id][IsImpression], expMap.m[id][IsClick], expMap.m[id][IsOrder], expMap.m[id][Total]
	controlIndicator.ImpressionCnt, controlIndicator.ClickCnt, controlIndicator.OrderCnt, controlIndicator.Total =
		expMap.m[controlEntity.ID][IsImpression], expMap.m[controlEntity.ID][IsClick], expMap.m[controlEntity.ID][IsOrder], expMap.m[controlEntity.ID][Total]
	// 释放实验统计结果的读锁
	expMap.RUnlock()

	// 停留量
	realRes = append(realRes, model.EffectRes{Param0: "停留量", Param1: strconv.Itoa(expIndicator.ImpressionCnt),
		Param2: strconv.Itoa(controlIndicator.ImpressionCnt), Param3: strconv.Itoa(expIndicator.ImpressionCnt - controlIndicator.ImpressionCnt), Param4: "-", Param5: "-"})

	// 点击量
	realRes = append(realRes, model.EffectRes{Param0: "点击量", Param1: strconv.Itoa(expIndicator.ClickCnt),
		Param2: strconv.Itoa(controlIndicator.ClickCnt), Param3: strconv.Itoa(expIndicator.ClickCnt - controlIndicator.ClickCnt), Param4: "-", Param5: "-"})

	// 下单量
	realRes = append(realRes, model.EffectRes{Param0: "下单量", Param1: strconv.Itoa(expIndicator.OrderCnt),
		Param2: strconv.Itoa(controlIndicator.OrderCnt), Param3: strconv.Itoa(expIndicator.OrderCnt - controlIndicator.OrderCnt), Param4: "-", Param5: "-"})

	// 流量总量
	realRes = append(realRes, model.EffectRes{Param0: "流量总量", Param1: strconv.Itoa(expIndicator.Total),
		Param2: strconv.Itoa(controlIndicator.Total), Param3: strconv.Itoa(expIndicator.Total - controlIndicator.Total), Param4: "-", Param5: "-"})

	expIndicator.Ctr, expIndicator.Cvr, controlIndicator.Ctr, controlIndicator.Cvr =
		utils.GetCR(expIndicator.ClickCnt, expIndicator.ImpressionCnt),
		utils.GetCR(expIndicator.OrderCnt, expIndicator.ClickCnt),
		utils.GetCR(controlIndicator.ClickCnt, controlIndicator.ImpressionCnt),
		utils.GetCR(controlIndicator.OrderCnt, controlIndicator.ClickCnt)

	// CTR
	realRes = append(realRes, model.EffectRes{Param0: "CTR", Param1: strconv.FormatFloat(expIndicator.Ctr*100, 'f', 2, 64) + "%",
		Param2: strconv.FormatFloat(controlIndicator.Ctr*100, 'f', 2, 64) + "%",
		Param3: strconv.FormatFloat(expIndicator.Ctr*100-controlIndicator.Ctr*100, 'f', 2, 64) + "%",
		Param4: strconv.FormatFloat(((expIndicator.Ctr-controlIndicator.Ctr)/controlIndicator.Ctr)*100, 'f', 2, 64) + "%",
		Param5: "-"})

	// CVR
	realRes = append(realRes, model.EffectRes{Param0: "CVR", Param1: strconv.FormatFloat(expIndicator.Cvr*100, 'f', 2, 64) + "%",
		Param2: strconv.FormatFloat(controlIndicator.Cvr*100, 'f', 2, 64) + "%",
		Param3: strconv.FormatFloat(expIndicator.Cvr*100-controlIndicator.Cvr*100, 'f', 2, 64) + "%",
		Param4: strconv.FormatFloat(((expIndicator.Cvr-controlIndicator.Cvr)/controlIndicator.Cvr)*100, 'f', 2, 64) + "%",
		Param5: "-"})

	expIndicator.CtrInterval, expIndicator.CvrInterval, controlIndicator.CtrInterval, controlIndicator.CvrInterval =
		utils.GetConfidenceIntervalByTwo(controlIndicator.Ctr, expIndicator.Ctr, controlIndicator.Total, expIndicator.Total),
		utils.GetConfidenceIntervalByTwo(controlIndicator.Cvr, expIndicator.Cvr, controlIndicator.Total, expIndicator.Total),
		utils.GetConfidenceInterval(controlIndicator.Ctr, controlIndicator.Total),
		utils.GetConfidenceInterval(controlIndicator.Cvr, controlIndicator.Total)

	// CTR置信区间
	realRes = append(realRes, model.EffectRes{Param0: "CTR置信区间", Param1: "[" + strconv.FormatFloat(expIndicator.CtrInterval[0], 'f', 2, 64) + "," +
		strconv.FormatFloat(expIndicator.CtrInterval[1], 'f', 2, 64) + "]",
		Param2: "[" + strconv.FormatFloat(controlIndicator.CtrInterval[0], 'f', 2, 64) + "," +
			strconv.FormatFloat(controlIndicator.CtrInterval[1], 'f', 2, 64) + "]",
		Param3: "-",
		Param4: "-",
		Param5: "-"})

	// CVR置信区间
	realRes = append(realRes, model.EffectRes{Param0: "CTR置信区间", Param1: "[" + strconv.FormatFloat(expIndicator.CvrInterval[0], 'f', 2, 64) + "," +
		strconv.FormatFloat(expIndicator.CvrInterval[1], 'f', 2, 64) + "]",
		Param2: "[" + strconv.FormatFloat(controlIndicator.CvrInterval[0], 'f', 2, 64) + "," +
			strconv.FormatFloat(controlIndicator.CvrInterval[1], 'f', 2, 64) + "]",
		Param3: "-",
		Param4: "-",
		Param5: "-"})

	return realRes, err
}
