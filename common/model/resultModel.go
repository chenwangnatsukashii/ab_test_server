package model

type ResultModel struct {
	PublishId    int  `json:"publish_id"`
	ExpId        int  `json:"exp_id"`
	IsImpression bool `json:"is_impression"`
	IsClick      bool `json:"is_click"`
	IsOrder      bool `json:"is_order"`
	UserId       int  `json:"user_id"`
}
