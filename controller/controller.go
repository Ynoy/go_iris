package controller

import (
	"github.com/kataras/iris"
	"lbps/service"
)

func DeviceInfoAction(ctx iris.Context) {
	path :=service.Query(ctx)

	ctx.JSON(iris.Map{"code": 1000, "data": path })
}

