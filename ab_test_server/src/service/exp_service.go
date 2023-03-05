package service

import (
	"line_china/ab_test_server/src/dao"
	"line_china/ab_test_server/src/model"
)

// AddExp 添加
func AddExp(model *model.Experiment) error {
	return dao.AddExp(model)
}

// DeleteExp 根据ID删除
func DeleteExp(id int) error {
	return dao.DeleteExp(id)
}

// OnLineExp 上线
func OnLineExp(id int) error {
	return dao.OnLineExp(id)
}

// OffLineExp 下线
func OffLineExp(id int) error {
	return dao.OffLineExp(id)
}

// SelectExp 根据Id查询
func SelectExp(id int) (*model.Experiment, error) {
	entity := new(model.Experiment)
	return dao.SelectExp(entity, id)
}

// SelectExps 查询所有
func SelectExps(pid int) ([]model.Experiment, error) {
	var models []model.Experiment
	return dao.SelectExps(models, pid)
}

// SelectAllExps 查询所有
func SelectAllExps() ([]model.Experiment, error) {
	var models []model.Experiment
	return dao.SelectAllExps(models)
}

// SelectOnLineExps 查询所有上线
func SelectOnLineExps(pid int) ([]model.Experiment, error) {
	var models []model.Experiment
	return dao.SelectOnLineExps(models, pid)
}

// UpdateExp 根据Id更新
func UpdateExp(model *model.Experiment) error {
	return dao.UpdateExp(model)
}
