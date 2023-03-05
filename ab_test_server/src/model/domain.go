package model

type Domain struct {
	Common
	Pid    int `json:"pid"`
	Weight int `json:"weight"`
}
