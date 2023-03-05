package service

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/dao"
	"line_china/ab_test_server/src/model"
	"time"
)

// AddLayer 添加
func AddLayer(model *model.Layer) error {
	model.State = constant.NormalState
	model.CreateTime = time.Now()
	model.UpdateTime = time.Now()
	return dao.AddLayer(model)
}

// DeleteLayer 根据ID删除
func DeleteLayer(id int) error {
	return dao.DeleteLayer(id)
}

// OnLineLayer 上线
func OnLineLayer(id int) error {
	return dao.OnLineLayer(id)
}

// OffLineLayer 下线
func OffLineLayer(id int) error {
	return dao.OffLineLayer(id)
}

// SelectLayer 根据Id查询
func SelectLayer(id int) (*model.Layer, error) {
	entity := new(model.Layer)
	return dao.SelectLayer(entity, id)
}

// SelectLayers 查询所有
func SelectLayers(pid int) ([]model.Layer, error) {
	var models []model.Layer
	return dao.SelectLayers(models, pid)
}

// SelectAllLayers 查询所有
func SelectAllLayers() ([]model.Layer, error) {
	var models []model.Layer
	return dao.SelectAllLayers(models)
}

// SelectOnLineLayers 查询所有上线
func SelectOnLineLayers(pid int) ([]model.Layer, error) {
	var models []model.Layer
	return dao.SelectOnLineLayers(models, pid)
}

// UpdateLayer 根据Id更新
func UpdateLayer(model *model.Layer) error {
	return dao.UpdateLayer(model)
}
