package v1

import "gogo/backend/app/service"

type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var (
	dashboardService = service.NewIDashboardService()
)
