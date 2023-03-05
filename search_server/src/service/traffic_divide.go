package service

import (
	"fmt"
	comModel "line_china/common/model"
	"line_china/search_server/src/utils"
	"math"
	"strconv"
)

// CalculateWeight 计算域权重
func CalculateWeight() comModel.Publish {
	publish = utils.ConfigData

	domains := publish.Domains
	var domainWeight []int
	for i := 0; i < len(domains); i++ {
		domainWeight = append(domainWeight, domains[i].Weight)
	}
	domainWeight = WeightNorm(domainWeight)
	for i := 0; i < len(domains); i++ {
		domains[i].Weight = domainWeight[i]
	}

	for _, domain := range domains {
		layers := domain.Layers
		for _, layer := range layers {
			exps := layer.Exps
			var expWeight []int
			for i := 0; i < len(exps); i++ {
				expWeight = append(expWeight, exps[i].Weight)
			}
			expWeight = WeightNorm(expWeight)
			for i := 0; i < len(exps); i++ {
				exps[i].Weight = expWeight[i]
			}
		}
	}

	return publish
}

// WeightNorm 根据输入比例计算归一化之后的比例
func WeightNorm(proportion []int) []int {
	var sum, total, hundred = 0, 0, 100
	for _, v := range proportion {
		total += v
	}

	num0, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(hundred)/float64(total)), 64) // 保留2位小数

	var res []int
	for i := 0; i < len(proportion)-1; i++ {
		res = append(res, int(math.Floor(num0*float64(proportion[i]))))
		sum += res[i]
	}

	return append(res, hundred-sum)
}
