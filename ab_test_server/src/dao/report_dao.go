package dao

import (
	"line_china/ab_test_server/src/common/constant"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
)

func AddReport(report model.Report) error {
	return utils.GormDb.Table(constant.ReportTable).Create(&report).Error
}
