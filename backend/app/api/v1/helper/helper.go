package helper

import (
	"github.com/gin-gonic/gin"
	"gogo/backend/app/dto"
	"net/http"
)

func SuccessWithData(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := dto.Response{
		Code: 200,
		Data: data,
	}
	ctx.JSON(http.StatusOK, res)
	ctx.Abort()
}
