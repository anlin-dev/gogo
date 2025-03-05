package router

import (
	"github.com/gin-gonic/gin"
	v1 "gogo/backend/app/api/v1"
)

type DashboardRouter struct{}

func (s *DashboardRouter) InitRouter(Router *gin.RouterGroup) {
	cmdRouter := Router.Group("dashboard")
	baseApi := v1.ApiGroupApp.BaseApi
	{
		cmdRouter.GET("/base/os", baseApi.LoadDashboardOsInfo)
	}
}
