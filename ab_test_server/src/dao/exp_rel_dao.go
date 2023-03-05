package dao

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
)

func AddExpRel(model *model.ExpRel) error {
	return utils.GormDb.Table(constant.ExpRelTable).Create(&model).Error
}

func SelectExpRel(entity model.ExpRel, pid int32) model.ExpRel {
	utils.GormDb.Table(constant.ExpRelTable).Where("pid=?", pid).Last(&entity)
	return entity
}
