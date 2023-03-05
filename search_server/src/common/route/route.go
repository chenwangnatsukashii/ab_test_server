package route

import (
	"github.com/gin-gonic/gin"
	"line_china/search_server/src/controller"
)

// PathRoute 接口路由
func PathRoute(r *gin.Engine) *gin.Engine {

	rootPath := r.Group("/line_china/search_server")
	{
		controller.SelectApps(rootPath)
		controller.ExpStartUp(rootPath)
		controller.ExpShutDown(rootPath)
	}

	return r
}
