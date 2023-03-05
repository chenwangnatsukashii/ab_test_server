package dao

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
)

func AddLayer(model *model.Layer) error {
	return utils.GormDb.Table(constant.LayerTable).Create(&model).Error
}

func DeleteLayer(id int) error {
	return utils.GormDb.Table(constant.LayerTable).Where("id=?", id).
		Update("state", constant.DeleteState).Error
}

func OnLineLayer(id int) error {
	return utils.GormDb.Table(constant.LayerTable).Where("id=?", id).
		Update("state", constant.OnLineState).Error
}

func OffLineLayer(id int) error {
	return utils.GormDb.Table(constant.LayerTable).Where("id=?", id).
		Update("state", constant.OffLineState).Error
}

func SelectLayer(entity *model.Layer, id int) (*model.Layer, error) {
	if err := utils.GormDb.Table(constant.LayerTable).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func SelectLayers(models []model.Layer, pid int) ([]model.Layer, error) {
	if err := utils.GormDb.Table(constant.LayerTable).Where("pid=?", pid).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func SelectAllLayers(models []model.Layer) ([]model.Layer, error) {
	if err := utils.GormDb.Table(constant.LayerTable).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func SelectOnLineLayers(models []model.Layer, pid int) ([]model.Layer, error) {
	if err := utils.GormDb.Table(constant.LayerTable).Where("pid=?", pid).
		Where("state=?", constant.OnLineState).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func UpdateLayer(model *model.Layer) error {
	return utils.GormDb.Table(constant.LayerTable).Where("id=?", model.ID).Updates(map[string]interface{}{
		"name":           model.Name,
		"default_config": model.DefaultConfig}).Error
}
