package utils

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
	"strconv"
)

const ZScore = 1.96

// GetCR 计算CTR或CVR
func GetCR(cnt int, total int) float64 {
	if total == 0 {
		return 0
	}
	CR, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", float64(cnt)/float64(total)), 64) // 保留4位小数
	return CR
}

// GetZScore 如果z大于1.9599639则拒绝原假设，说明实验组效果好
func GetZScore(controlHit, controlNum, expHit, expNum int) error {
	var err error
	var cra, crb, r float64
	cra, err = strconv.ParseFloat(fmt.Sprintf("%.4f", float64(controlHit)/float64(controlNum)), 64) // 保留4位小数
	if err != nil {
		return err
	}
	crb, err = strconv.ParseFloat(fmt.Sprintf("%.4f", float64(expHit)/float64(expNum)), 64) // 保留4位小数
	if err != nil {
		return err
	}
	r, err = strconv.ParseFloat(fmt.Sprintf("%.4f", float64(controlHit+expHit)/float64(controlNum+expNum)), 64)
	// 计算检验统计量Z
	z := (crb - cra) / math.Sqrt(r*(1-r)*(float64(1)/float64(controlNum)+float64(1)/float64(expNum)))
	fmt.Println("检验统计量z:", z)

	res := GetConfidenceIntervalByTwo(cra, crb, controlNum, expNum)
	fmt.Println(res[0], "   ", res[1])
	return nil
}

// TwoSidedPValue 获取p-value
func TwoSidedPValue(r float64, n float64) float64 {
	t := distuv.StudentsT{Sigma: 1, Nu: n - 2}
	pVal := t.CDF(-math.Abs(r))
	return pVal
}

// GetConfidenceInterval 获取单个的置信区间
func GetConfidenceInterval(r float64, cnt int) [2]float64 {
	delta := ZScore * math.Sqrt(r*(1-r)/float64(cnt))

	res := [2]float64{}
	res[0] = r - delta
	res[1] = r + delta
	return res
}

// GetConfidenceIntervalByTwo 计算对照组和实验组的置信区间
func GetConfidenceIntervalByTwo(cra, crb float64, sa, sb int) [2]float64 {
	cru := crb - cra
	delta := ZScore * math.Sqrt(cra*(1-cra)/float64(sa)+crb*(1-crb)/float64(sb))

	res := [2]float64{}
	res[0] = cru - delta
	res[1] = cru + delta
	return res
}
