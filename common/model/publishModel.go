package model

type Publish struct {
	AppId     int            `json:"app_id"`
	Name      string         `json:"name"`
	PublishId int            `json:"publish_id"`
	Domains   []PublicDomain `json:"domains"`
}

type PublicDomain struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Weight int     `json:"weight"`
	Layers []Layer `json:"Layers"`
}

type Layer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Exps []Exp  `json:"exps"`
}

type Exp struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Weight    int    `json:"weight"`
	ExpConfig JSON   `json:"exp_config"`
	IsControl int    `json:"is_control"`
}
