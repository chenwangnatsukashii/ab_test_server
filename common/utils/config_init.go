package utils

import "line_china/ab_test_server/src/model"

var err error
var Config *model.Config

func ConfigInit() error {
	Config, err = model.NewConfig().ReadConfig()
	return err
}
