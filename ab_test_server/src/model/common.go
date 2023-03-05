package model

import (
	"time"
)

type Common struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	State      int       `json:"state"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
