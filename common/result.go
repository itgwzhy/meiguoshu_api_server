package common

import (
	"github.com/kataras/iris"
)

func Responese(ctx iris.Context, code int, data interface{}, message string) {
	dataMap := iris.Map{
		"code": code,
		"data": data,
		"msg":  message,
	}
	ctx.JSON(dataMap)
}
