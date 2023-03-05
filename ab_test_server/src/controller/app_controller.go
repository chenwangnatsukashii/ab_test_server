package controller

import (
	"github.com/gin-gonic/gin"
	"line_china/ab_test_server/src/common/response"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/service"
	"strconv"
)

type AppController struct {
}

// AddApp 添加
func AddApp(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().POST("", controller.AddApp)
}

// SelectApp 根据ID查询
func SelectApp(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().GET("/:id", controller.SelectApp)
}

// SelectApps 查询所有
func SelectApps(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().GET("", controller.SelectApps)
}

// DeleteApp 根据ID删除
func DeleteApp(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().DELETE("/:id", controller.DeleteApp)
}

// OnLineApp 上线
func OnLineApp(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().PUT("/online/:id", controller.OnLineApp)
}

// OffLineApp 下线
func OffLineApp(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().PUT("/offline/:id", controller.OffLineApp)
}

// UpdateApp 更新
func UpdateApp(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().PUT("", controller.UpdateApp)
}

func (c AppController) AddApp(ctx *gin.Context) {
	var entity model.Application
	var err error
	if err = ctx.ShouldBind(&entity); err == nil {
		if err = service.AddApp(&entity); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) SelectApp(ctx *gin.Context) {
	var id int
	var err error
	var app *model.Application
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		app, err = service.SelectApp(id)
		if err == nil {
			response.Success(ctx, app)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) SelectApps(ctx *gin.Context) {
	models, err := service.SelectApps()
	if err == nil {
		response.Success(ctx, models)
		return
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) DeleteApp(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.DeleteApp(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) OnLineApp(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.OnLineApp(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) OffLineApp(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.OffLineApp(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) UpdateApp(ctx *gin.Context) {
	var entity model.Application
	var err error
	if err = ctx.ShouldBind(&entity); err == nil {
		if err = service.UpdateApp(&entity); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}
