package dao

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
)

func AddApp(model *model.Application) error {
	return utils.GormDb.Table(constant.AppTable).Create(&model).Error
}

func SelectApp(entity *model.Application, id int) (*model.Application, error) {
	if err := utils.GormDb.Table(constant.AppTable).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func SelectApps(models []model.Application) ([]model.Application, error) {
	if err := utils.GormDb.Table(constant.AppTable).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func DeleteApp(id int) error {
	return utils.GormDb.Table(constant.AppTable).Where("id=?", id).
		Update("state", constant.DeleteState).Error
}

func OnLineApp(id int) error {
	return utils.GormDb.Table(constant.AppTable).Where("id=?", id).
		Update("state", constant.OnLineState).Error
}

func OffLineApp(id int) error {
	return utils.GormDb.Table(constant.AppTable).Where("id=?", id).
		Update("state", constant.OffLineState).Error
}

func UpdateApp(model *model.Application) error {
	return utils.GormDb.Table(constant.AppTable).Where("id=?", model.ID).
		Update("name", model.Name).Error
}
