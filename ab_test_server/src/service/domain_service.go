package service

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/dao"
	"line_china/ab_test_server/src/model"
	"time"
)

// AddDomain 添加
func AddDomain(model *model.Domain) error {
	model.State = constant.NormalState
	model.CreateTime = time.Now()
	model.UpdateTime = time.Now()
	return dao.AddDomain(model)
}

// DeleteDomain 根据ID删除
func DeleteDomain(id int) error {
	return dao.DeleteDomain(id)
}

// OnLineDomain 上线
func OnLineDomain(id int) error {
	return dao.OnLineDomain(id)
}

// OffLineDomain 下线
func OffLineDomain(id int) error {
	return dao.OffLineDomain(id)
}

// SelectDomain 根据Id查询
func SelectDomain(id int) (*model.Domain, error) {
	entity := new(model.Domain)
	return dao.SelectDomain(entity, id)
}

// SelectDomains 查询所有
func SelectDomains(pid int) ([]model.Domain, error) {
	var models []model.Domain
	return dao.SelectDomains(models, pid)
}

// SelectAllDomains 查询所有
func SelectAllDomains() ([]model.Domain, error) {
	var models []model.Domain
	return dao.SelectAllDomains(models)
}

// SelectOnLineDomains 查询所有上线的
func SelectOnLineDomains(pid int) ([]model.Domain, error) {
	var models []model.Domain
	return dao.SelectOnLineDomains(models, pid)
}

// UpdateDomain 根据Id更新
func UpdateDomain(entity *model.Domain) error {
	return dao.UpdateDomain(entity)
}
