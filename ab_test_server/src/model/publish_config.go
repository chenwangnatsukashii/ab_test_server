package model

import (
	common "line_china/common/model"
	"time"
)

type PublishConfig struct {
	ID         int         `json:"id"`
	ExpConfig  common.JSON `json:"exp_config"`
	CreateTime time.Time   `json:"create_time"`
}
