package controller

import (
	"github.com/gin-gonic/gin"
	"line_china/ab_test_server/src/common/response"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/service"
	"strconv"
)

type DomainController struct {
}

// AddDomain 添加
func AddDomain(grp *gin.RouterGroup) {
	controller := DomainController{}
	grp.Use().POST("", controller.AddDomain)
}

// SelectDomain 根据ID查询
func SelectDomain(grp *gin.RouterGroup) {
	controller := DomainController{}
	grp.Use().GET("/:id", controller.SelectDomain)
}

// SelectDomains 根据ID查询所有
func SelectDomains(grp *gin.RouterGroup) {
	controller := DomainController{}
	grp.Use().GET("/pid/:pid", controller.SelectDomains)
}

// SelectAllDomains 查询所有
func SelectAllDomains(grp *gin.RouterGroup) {
	controller := DomainController{}
	grp.Use().GET("", controller.SelectAllDomains)
}

// DeleteDomain 根据ID删除
func DeleteDomain(grp *gin.RouterGroup) {
	controller := DomainController{}
	grp.Use().DELETE("/:id", controller.DeleteDomain)
}

// OnLineDomain 上线
func OnLineDomain(grp *gin.RouterGroup) {
	controller := DomainController{}
	grp.Use().PUT("/online/:id", controller.OnLineDomain)
}

// OffLineDomain 下线
func OffLineDomain(grp *gin.RouterGroup) {
	controller := DomainController{}
	grp.Use().PUT("/offline/:id", controller.OffLineDomain)
}

// UpdateDomain 更新
func UpdateDomain(grp *gin.RouterGroup) {
	controller := DomainController{}
	grp.Use().PUT("", controller.UpdateDomain)
}

func (c DomainController) AddDomain(ctx *gin.Context) {
	var entity model.Domain
	var err error
	if err = ctx.ShouldBind(&entity); err == nil {
		if err = service.AddDomain(&entity); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c DomainController) SelectDomain(ctx *gin.Context) {
	var id int
	var err error
	var domain *model.Domain
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		domain, err = service.SelectDomain(id)
		if err == nil {
			response.Success(ctx, domain)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c DomainController) SelectAllDomains(ctx *gin.Context) {
	models, err := service.SelectAllDomains()
	if err == nil {
		response.Success(ctx, models)
		return
	}
	response.Failure(ctx, err.Error())
}

func (c DomainController) SelectDomains(ctx *gin.Context) {
	var pid int
	var err error
	var models []model.Domain
	pid, err = strconv.Atoi(ctx.Param("pid"))
	if err == nil {
		models, err = service.SelectDomains(pid)
		if err == nil {
			response.Success(ctx, models)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c DomainController) DeleteDomain(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.DeleteDomain(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c DomainController) OnLineDomain(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.OnLineDomain(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c DomainController) OffLineDomain(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err == nil {
		if err = service.OffLineDomain(id); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}

func (c DomainController) UpdateDomain(ctx *gin.Context) {
	var entity model.Domain
	var err error
	if err = ctx.ShouldBind(&entity); err == nil {
		if err = service.UpdateDomain(&entity); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Failure(ctx, err.Error())
}
