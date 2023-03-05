package controller

import (
	"github.com/gin-gonic/gin"
	"line_china/common/constant"
	"line_china/search_server/src/common/http"
	"line_china/search_server/src/common/response"
	"line_china/search_server/src/service"
	"net/http"
	"strconv"
	_ "strings"
	"time"
)

type SdkController struct {
}

// SelectApps 获取app列表
func SelectApps(grp *gin.RouterGroup) {
	controller := SdkController{}
	grp.Use().GET("/apps", controller.SelectApps)
}

// ExpStartUp 开始实验
func ExpStartUp(grp *gin.RouterGroup) {
	controller := SdkController{}
	grp.Use().GET("/startup/:number", controller.ExpStartUp)
}

// ExpShutDown 结束实验
func ExpShutDown(grp *gin.RouterGroup) {
	controller := SdkController{}
	grp.Use().GET("/shutdown/:id", controller.ExpShutDown)
}

func (c SdkController) ExpStartUp(ctx *gin.Context) {
	number, err := strconv.Atoi(ctx.Param("number"))
	if err == nil {
		go service.StartUp(number)
		response.Success(ctx, constant.SuccessMsg, nil)
		return
	}
	response.Failure(ctx, constant.ErrorMsg, err.Error())
}

func (c SdkController) ExpShutDown(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err == nil {
		go service.ShutDown(id)
		response.Success(ctx, constant.SuccessMsg, nil)
		return
	}
	response.Failure(ctx, constant.ErrorMsg, err.Error())
}

func (c SdkController) SelectApps(ctx *gin.Context) {

	classDetailMap, costTime := line_http.HttpGet("/app")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":       "请求成功",
		"time":      time.Now().Unix(),
		"data":      classDetailMap["data"],
		"cost_time": costTime,
	})

}
