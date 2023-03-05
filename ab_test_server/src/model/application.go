package model

type Application struct {
	Common
}

func (Application) TableName() string {
	return "application"
}
