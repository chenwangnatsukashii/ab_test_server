package utils

import (
	"line_china/ab_test_server/src/model"
	"log"
)

var Config *model.Config

func ConfigInit() {
	var err error
	if Config, err = model.NewConfig().ReadConfig(); err != nil {
		log.Println(err)
	}
}
