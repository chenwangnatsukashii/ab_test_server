package controller

import (
	"github.com/gin-gonic/gin"
	"line_china/ab_test_server/src/common/response"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/service"
	"strconv"
)

type IndicatorController struct {
}

// GetIndicatorById 根据实验Id获取指标
func GetIndicatorById(grp *gin.RouterGroup) {
	controller := ExpController{}
	grp.Use().GET("/expId/:id", controller.GetIndicatorById)
}

// GetIndicatorByRange 根据实验Id获取指标
// TODO
func GetIndicatorByRange(grp *gin.RouterGroup) {
	controller := ExpController{}
	grp.Use().GET("/range/:start", controller.GetIndicatorById)
}

func (c ExpController) GetIndicatorById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	var res []model.EffectRes
	if err == nil {
		if res, err = service.GetIndicatorById(id); err == nil {
			response.Success(ctx, res)
			return
		}
	}
	response.Failure(ctx, err.Error())
}
