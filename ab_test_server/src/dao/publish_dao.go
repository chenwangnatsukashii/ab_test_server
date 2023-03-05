package dao

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
)

func AddPublish(entity model.PublishConfig) (error, int) {
	if err := utils.GormDb.Table(constant.PublishTable).Create(&entity).Error; err != nil {
		return err, 0
	}
	return nil, entity.ID
}
