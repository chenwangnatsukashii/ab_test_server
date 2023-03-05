package controller

import (
	"github.com/gin-gonic/gin"
	"line_china/ab_test_server/src/common/response"
	"line_china/ab_test_server/src/service"
	"strconv"
)

type PublishController struct {
}

// PublishAll 根据ID发布
func PublishAll(grp *gin.RouterGroup) {
	controller := PublishController{}
	grp.Use().POST("/:id", controller.PublishAll)
}

// GetPublishReview 根据ID获取预览
func GetPublishReview(grp *gin.RouterGroup) {
	controller := PublishController{}
	grp.Use().GET("/review/:id", controller.GetPublishReview)
}

func (c PublishController) PublishAll(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err == nil {
		var res map[string]interface{}
		res, err = service.PublishAll(id)
		if err == nil {
			response.Success(ctx, res)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c PublishController) GetPublishReview(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err == nil {
		var res map[string]interface{}
		res, err = service.GetPublishReview(id)
		if err == nil {
			response.Success(ctx, res)
			return
		}
	}
	response.Failure(ctx, err.Error())
}
