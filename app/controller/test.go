package controller

import (
	"gin-project/app/lib/response"
	"gin-project/app/openapi"

	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}

	commonResponse := openapi.CommonResponse{
		Status:  openapi.RESPONSESTATUS_SUCCESS,
		Message: "holle~",
	}

	resp.Success(commonResponse)
}

// test相关api
var TestApi = &openapi.TestAPI{
	Say: SayHello,
}
