package model

import common "line_china/common/model"

type Experiment struct {
	Common
	Pid       int         `json:"pid"`
	Weight    int         `json:"weight"`
	ExpConfig common.JSON `json:"exp_config"` // gorm:"-" 用于忽略这个字段
	IsControl int         `json:"is_control"`
	LayerName string      `json:"layer_name"`
}
