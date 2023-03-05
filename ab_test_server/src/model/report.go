package model

import (
	"time"
)

type Report struct {
	ID            int       `json:"id"`
	PublishId     int       `json:"publish_id"`
	ExpId         int       `json:"exp_id"`
	ImpressionCnt int       `json:"impression_cnt"`
	ClickCnt      int       `json:"click_cnt"`
	OrderCnt      int       `json:"order_cnt"`
	Total         int       `json:"total"`
	DateInfo      time.Time `json:"date_info"`
}
