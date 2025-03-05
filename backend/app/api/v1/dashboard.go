package v1

import (
	"github.com/gin-gonic/gin"
	"gogo/backend/app/api/v1/helper"
)

func (b *BaseApi) LoadDashboardOsInfo(c *gin.Context) {

	data, err := dashboardService.LoadOsInfo()
	if err != nil {
		return
	}
	helper.SuccessWithData(c, data)
}
