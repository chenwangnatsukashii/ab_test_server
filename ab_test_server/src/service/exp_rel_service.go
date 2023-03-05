package service

import (
	"line_china/ab_test_server/src/dao"
	"line_china/ab_test_server/src/model"
	"time"
)

// AddExpRel 添加
func AddExpRel(model *model.ExpRel) error {
	model.CreateTime = time.Now()
	return dao.AddExpRel(model)
}

// SelectExpRel 根据Id查询
func SelectExpRel(pid int32) model.ExpRel {
	entity := model.ExpRel{}
	return dao.SelectExpRel(entity, pid)
}
