package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(ctx *gin.Context,message string, code int, status string,data interface{}){
	meta := Meta{
		Message: message,
		Code: code,
		Status: status,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	ctx.JSON(code,jsonResponse)
}

func APIErrorResponse(ctx *gin.Context,message string,code int,status string,errDetails interface{}){
	meta := Meta{
		Message: message,
		Code: code,
		Status: status,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: errDetails,
	}
	ctx.JSON(code,jsonResponse)
}