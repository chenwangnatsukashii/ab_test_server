package model

type Indicator struct {
	ExpName       string     `json:"exp_name"`
	ImpressionCnt int        `json:"impression_cnt"`
	ClickCnt      int        `json:"click_cnt"`
	OrderCnt      int        `json:"order_cnt"`
	Total         int        `json:"total"`
	Ctr           float64    `json:"ctr"`
	Cvr           float64    `json:"cvr"`
	CtrInterval   [2]float64 `json:"ctrInterval"`
	CvrInterval   [2]float64 `json:"cvrInterval"`
}
