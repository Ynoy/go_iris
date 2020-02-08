package web

import (
	"github.com/kataras/iris"
	"lbps/controller"
)

// 定义500错误处理函数
func err500(ctx iris.Context) {
	ctx.WriteString("CUSTOM 500 ERROR")
}

// 定义404错误处理函数
func err404(ctx iris.Context) {
	ctx.WriteString("CUSTOM 404 ERROR")
}

// 注入路由
func InnerRoute(app *iris.Application) {
	app.OnErrorCode(iris.StatusInternalServerError, err500)
	app.OnErrorCode(iris.StatusNotFound, err404)
	app.Get("/ping", func(ctx iris.Context) { ctx.WriteString("pong") })

	// 导出Excel数据
	app.Post("/export/excel", controller.DeviceInfoAction)
}
