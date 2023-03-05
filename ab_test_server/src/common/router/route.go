package router

import (
	"github.com/gin-gonic/gin"
	"line_china/ab_test_server/src/controller"
)

// PathRoute 路由
func PathRoute(r *gin.Engine) *gin.Engine {

	rootPath := r.Group("/line_china/abtest")
	{
		appPath := rootPath.Group("/app")
		{
			controller.SelectApp(appPath)
			controller.DeleteApp(appPath)
			controller.SelectApps(appPath)
			controller.AddApp(appPath)
			controller.UpdateApp(appPath)
			controller.OnLineApp(appPath)
			controller.OffLineApp(appPath)
		}

		domainPath := rootPath.Group("/domain")
		{
			controller.SelectDomain(domainPath)
			controller.DeleteDomain(domainPath)
			controller.SelectDomains(domainPath)
			controller.SelectAllDomains(domainPath)
			controller.AddDomain(domainPath)
			controller.UpdateDomain(domainPath)
			controller.OnLineDomain(domainPath)
			controller.OffLineDomain(domainPath)
		}

		layerPath := rootPath.Group("/layer")
		{
			controller.SelectLayer(layerPath)
			controller.DeleteLayer(layerPath)
			controller.SelectLayers(layerPath)
			controller.SelectAllLayers(layerPath)
			controller.AddLayer(layerPath)
			controller.UpdateLayer(layerPath)
			controller.OnLineLayer(layerPath)
			controller.OffLineLayer(layerPath)
		}

		expPath := rootPath.Group("/experiment")
		{
			controller.SelectExp(expPath)
			controller.DeleteExp(expPath)
			controller.SelectExps(expPath)
			controller.SelectAllExps(expPath)
			controller.AddExp(expPath)
			controller.UpdateExp(expPath)
			controller.OnLineExp(expPath)
			controller.OffLineExp(expPath)
		}

		publishPath := rootPath.Group("/publish")
		{
			controller.PublishAll(publishPath)
			controller.GetPublishReview(publishPath)
		}

		indicatorPath := rootPath.Group("/indicator")
		{
			controller.GetIndicatorById(indicatorPath)
		}

	}

	return r
}
