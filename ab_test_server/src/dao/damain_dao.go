package dao

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
)

func AddDomain(model *model.Domain) error {
	return utils.GormDb.Table(constant.DomainTable).Create(&model).Error
}

func DeleteDomain(id int) error {
	return utils.GormDb.Table(constant.DomainTable).Where("id=?", id).
		Update("state", constant.DeleteState).Error
}

func OnLineDomain(id int) error {
	return utils.GormDb.Table(constant.DomainTable).Where("id=?", id).
		Update("state", constant.OnLineState).Error
}

func OffLineDomain(id int) error {
	return utils.GormDb.Table(constant.DomainTable).Where("id=?", id).
		Update("state", constant.OffLineState).Error
}

func SelectDomain(entity *model.Domain, id int) (*model.Domain, error) {
	if err := utils.GormDb.Table(constant.DomainTable).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func SelectDomains(models []model.Domain, pid int) ([]model.Domain, error) {
	if err := utils.GormDb.Table(constant.DomainTable).Where("pid=?", pid).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func SelectAllDomains(models []model.Domain) ([]model.Domain, error) {
	if err := utils.GormDb.Table(constant.DomainTable).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func SelectOnLineDomains(models []model.Domain, pid int) ([]model.Domain, error) {
	if err := utils.GormDb.Table(constant.DomainTable).Where("pid=?", pid).
		Where("state=?", constant.OnLineState).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func UpdateDomain(entity *model.Domain) error {
	return utils.GormDb.Table(constant.DomainTable).Where("id=?", entity.Common.ID).Updates(map[string]interface{}{
		"name":   entity.Name,
		"weight": entity.Weight}).Error
}
