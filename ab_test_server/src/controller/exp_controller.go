package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"line_china/ab_test_server/src/common/response"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/service"
	"line_china/ab_test_server/src/utils"
	"strconv"
)

type ExpController struct {
}

// AddExp 添加
func AddExp(grp *gin.RouterGroup) {
	controller := ExpController{}
	grp.Use().POST("", controller.AddExp)
}

// SelectExp 根据ID查询
func SelectExp(grp *gin.RouterGroup) {
	controller := ExpController{}
	grp.Use().GET("/:id", controller.SelectExp)
}

// SelectExps 根据ID查询所有
func SelectExps(grp *gin.RouterGroup) {
	controller := ExpController{}
	grp.Use().GET("/pid/:pid", controller.SelectExps)
}

// SelectAllExps 查询所有
func SelectAllExps(grp *gin.RouterGroup) {
	controller := ExpController{}
	grp.Use().GET("", controller.SelectAllExps)
}

// DeleteExp 根据ID删除
func DeleteExp(grp *gin.RouterGroup) {
	controller := ExpController{}
	grp.Use().DELETE("/:id", controller.DeleteExp)
}

// OnLineExp 上线
func OnLineExp(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().PUT("/online/:id", controller.OnLineExp)
}

// OffLineExp 下线
func OffLineExp(grp *gin.RouterGroup) {
	controller := AppController{}
	grp.Use().PUT("/offline/:id", controller.OffLineExp)
}

// UpdateExp 更新
func UpdateExp(grp *gin.RouterGroup) {
	controller := ExpController{}
	grp.Use().PUT("", controller.UpdateExp)
}

func (c ExpController) AddExp(ctx *gin.Context) {
	var entity model.Experiment
	var err error
	if err = ctx.ShouldBind(&entity); err == nil {
		err = utils.GormDb.Transaction(func(tx *gorm.DB) error {
			if err = service.AddExp(&entity); err == nil {
				if err = service.AddExpRel(&model.ExpRel{Pid: entity.ID, ExpConfig: entity.ExpConfig}); err == nil {
					return nil
				}
			}
			return err
		})
	}

	if err == nil {
		response.Success(ctx, nil)
		return
	}
	response.Failure(ctx, err.Error())
}

func (c ExpController) SelectExp(ctx *gin.Context) {
	var id int
	var err error
	var exp *model.Experiment
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		exp, err = service.SelectExp(id)
		if err == nil {
			response.Success(ctx, exp)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c ExpController) SelectAllExps(ctx *gin.Context) {
	models, err := service.SelectAllExps()
	if err == nil {
		response.Success(ctx, models)
		return
	}
	response.Failure(ctx, err.Error())
}

func (c ExpController) SelectExps(ctx *gin.Context) {
	var pid int
	var err error
	var models []model.Experiment
	pid, err = strconv.Atoi(ctx.Param("pid"))
	if err == nil {
		models, err = service.SelectExps(pid)
		if err == nil {
			response.Success(ctx, models)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c ExpController) DeleteExp(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.DeleteExp(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) OnLineExp(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.OnLineExp(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c AppController) OffLineExp(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.OffLineExp(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c ExpController) UpdateExp(ctx *gin.Context) {
	var entity model.Experiment
	var err error
	if err = ctx.ShouldBind(&entity); err == nil {
		err = utils.GormDb.Transaction(func(tx *gorm.DB) error {
			if err = service.UpdateExp(&entity); err == nil {
				if err = service.AddExpRel(&model.ExpRel{Pid: entity.ID, ExpConfig: entity.ExpConfig}); err == nil {
					response.Success(ctx, nil)
					return nil
				}
			}
			return err
		})
	}
	response.Failure(ctx, err.Error())
}
