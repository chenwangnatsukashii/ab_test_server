package model

import (
	common "line_china/common/model"
	"time"
)

type ExpRel struct {
	ID         int         `json:"id"`
	Pid        int         `json:"pid"`
	ExpConfig  common.JSON `json:"exp_config"`
	CreateTime time.Time   `json:"create_time"`
}
