package service

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/dao"
	"line_china/ab_test_server/src/model"
	"time"
)

// AddApp 添加
func AddApp(model *model.Application) error {
	model.State = constant.NormalState
	model.CreateTime = time.Now()
	model.UpdateTime = time.Now()
	return dao.AddApp(model)
}

// SelectApp 根据Id查询
func SelectApp(id int) (*model.Application, error) {
	entity := new(model.Application)
	return dao.SelectApp(entity, id)
}

// SelectApps 查询所有
func SelectApps() ([]model.Application, error) {
	var models []model.Application
	return dao.SelectApps(models)
}

// DeleteApp 根据ID删除
func DeleteApp(id int) error {
	return dao.DeleteApp(id)
}

// OnLineApp 上线
func OnLineApp(id int) error {
	return dao.OnLineApp(id)
}

// OffLineApp 下线
func OffLineApp(id int) error {
	return dao.OffLineApp(id)
}

// UpdateApp 根据Id更新
func UpdateApp(model *model.Application) error {
	return dao.UpdateApp(model)
}
