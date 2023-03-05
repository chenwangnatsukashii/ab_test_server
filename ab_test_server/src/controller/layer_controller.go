package controller

import (
	"github.com/gin-gonic/gin"
	"line_china/ab_test_server/src/common/response"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/service"
	"strconv"
)

type LayerController struct {
}

// AddLayer 添加
func AddLayer(grp *gin.RouterGroup) {
	controller := LayerController{}
	grp.Use().POST("", controller.AddLayer)
}

// SelectLayer 根据ID查询
func SelectLayer(grp *gin.RouterGroup) {
	controller := LayerController{}
	grp.Use().GET("/:id", controller.SelectLayer)
}

// SelectLayers 根据ID查询所有
func SelectLayers(grp *gin.RouterGroup) {
	controller := LayerController{}
	grp.Use().GET("/pid/:pid", controller.SelectLayers)
}

// SelectAllLayers 查询所有
func SelectAllLayers(grp *gin.RouterGroup) {
	controller := LayerController{}
	grp.Use().GET("", controller.SelectAllLayers)
}

// DeleteLayer 根据ID删除
func DeleteLayer(grp *gin.RouterGroup) {
	controller := LayerController{}
	grp.Use().DELETE("/:id", controller.DeleteLayer)
}

// OnLineLayer 上线
func OnLineLayer(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().PUT("/online/:id", controller.OnLineLayer)
}

// OffLineLayer 下线
func OffLineLayer(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().PUT("/offline/:id", controller.OffLineLayer)
}

// UpdateLayer 更新
func UpdateLayer(grp *gin.RouterGroup) {
	controller := LayerController{}
	grp.Use().PUT("", controller.UpdateLayer)
}

func (c LayerController) AddLayer(ctx *gin.Context) {
	var entity model.Layer
	var err error
	if err = ctx.ShouldBind(&entity); err == nil {
		if err = service.AddLayer(&entity); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c LayerController) SelectLayer(ctx *gin.Context) {
	var id int
	var err error
	var layer *model.Layer
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		layer, err = service.SelectLayer(id)
		if err == nil {
			response.Success(ctx, layer)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c LayerController) SelectAllLayers(ctx *gin.Context) {
	models, err := service.SelectAllLayers()
	if err == nil {
		response.Success(ctx, models)
		return
	}
	response.Failure(ctx, err.Error())
}

func (c LayerController) SelectLayers(ctx *gin.Context) {
	var pid int
	var err error
	var models []model.Layer
	pid, err = strconv.Atoi(ctx.Param("pid"))
	if err == nil {
		models, err = service.SelectLayers(pid)
		if err == nil {
			response.Success(ctx, models)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c LayerController) DeleteLayer(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.DeleteLayer(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) OnLineLayer(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.OnLineLayer(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) OffLineLayer(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.OffLineLayer(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c LayerController) UpdateLayer(ctx *gin.Context) {
	var entity model.Layer
	var err error
	if err = ctx.ShouldBind(&entity); err == nil {
		if err = service.UpdateLayer(&entity); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, nil)
}
