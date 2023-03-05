package dao

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
	"time"
)

func AddExp(model *model.Experiment) error {
	return utils.GormDb.Table(constant.ExpTable).Create(map[string]interface{}{
		"pid":         model.Pid,
		"name":        model.Name,
		"state":       constant.NormalState,
		"weight":      model.Weight,
		"is_control":  model.IsControl,
		"create_time": time.Now(),
		"update_time": time.Now()}).Error
}

func DeleteExp(id int) error {
	return utils.GormDb.Table(constant.ExpTable).Where("id=?", id).
		Update("state", constant.DeleteState).Error
}

func OnLineExp(id int) error {
	return utils.GormDb.Table(constant.ExpTable).Where("id=?", id).
		Update("state", constant.OnLineState).Error
}

func OffLineExp(id int) error {
	return utils.GormDb.Table(constant.ExpTable).Where("id=?", id).
		Update("state", constant.OffLineState).Error
}

func SelectExp(entity *model.Experiment, id int) (*model.Experiment, error) {
	if err := utils.GormDb.Table(constant.ExpTable).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func SelectExps(models []model.Experiment, pid int) ([]model.Experiment, error) {
	if err := utils.GormDb.Table(constant.ExpTable).Select("experiment.*, layer.name as layer_name").
		Joins("JOIN layer ON experiment.pid = layer.id").
		Where("experiment.pid=?", pid).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func SelectAllExps(models []model.Experiment) ([]model.Experiment, error) {
	if err := utils.GormDb.Table(constant.ExpTable).Select("experiment.*, layer.name as layer_name").
		Joins("JOIN layer ON experiment.pid = layer.id").Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func SelectOnLineExps(models []model.Experiment, pid int) ([]model.Experiment, error) {
	if err := utils.GormDb.Table(constant.ExpTable).Where("pid=?", pid).
		Where("state=?", constant.OnLineState).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func UpdateExp(model *model.Experiment) error {
	return utils.GormDb.Table(constant.ExpTable).Where("id=?", model.ID).Updates(map[string]interface{}{
		"name":       model.Name,
		"weight":     model.Weight,
		"is_control": model.IsControl}).Error
}

func GetControl(pid int) (*model.Experiment, error) {
	var entity *model.Experiment

	if err := utils.GormDb.Table(constant.ExpTable).Where("is_control = ?", 1).
		Where("pid = ?", pid).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
