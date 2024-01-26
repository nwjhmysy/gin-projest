package controller

import (
	"gin-project/app/lib/response"
	"gin-project/app/openapi"

	"github.com/gin-gonic/gin"
)

type TestAPI struct {
	openapi.TestAPI
}

// 重写 test 标签的方法
func (api *TestAPI) SayHello(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}

	commonResponse := openapi.CommonResponse{
		Status:  openapi.RESPONSESTATUS_SUCCESS,
		Message: "holle~",
	}

	resp.Success(commonResponse)
}

var TestApi TestAPI
