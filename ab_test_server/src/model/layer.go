package model

import (
	common "line_china/common/model"
)

type Layer struct {
	Common
	Pid           int         `json:"pid"`
	DefaultConfig common.JSON `json:"default_config"`
}
