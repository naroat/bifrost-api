package system

import (
	v1 "bifrost/server/api/v1"
	"github.com/gin-gonic/gin"
)

type EcApiRouter struct{}

func (s *EcApiRouter) InitEcApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	router := Router.Group("go/ecapi")
	controller := v1.ApiGroupApp.SystemApiGroup.EcApi
	{
		router.GET("ele", controller.Ele)
	}
	return router
}
