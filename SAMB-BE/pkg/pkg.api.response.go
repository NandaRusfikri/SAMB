package pkg

import (
	"SAMB-BE/schemas"
	"github.com/gin-gonic/gin"
)

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Count int64, Data interface{}) {

	jsonResponse := schemas.SchemaResponses{
		Code:    StatusCode,
		Count:   Count,
		Message: Message,
		Data:    Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}
