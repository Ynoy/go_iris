package web

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"lbps/utility/helper"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func RunIris(port int) {
	app := iris.New()

	app.Use(NewRecoverMdw())
	app.Use(NewAccessLogMdw())

	// 优雅的关闭程序
	serverWG := new(sync.WaitGroup)
	defer serverWG.Wait()

	iris.RegisterOnInterrupt(func() {
		serverWG.Add(1)
		defer serverWG.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
		defer cancel()

		app.Shutdown(ctx)
	})

	// 注册路由
	InnerRoute(app)

	// server配置
	c := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           true,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: true,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "2020-01-31 10:04:05",
		Charset:                           "UTF-8",
		IgnoreServerErrors:                []string{iris.ErrServerClosed.Error()},
		RemoteAddrHeaders:                 map[string]bool{"X-Real-Ip": true, "X-Forwarded-For": true},
	})

	app.Run(iris.Addr(":"+strconv.Itoa(port)), c)
}

// 统一异常处理
func NewRecoverMdw() iris.Handler {
	return func(ctx iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("At Request: %d %s %s %s\n", ctx.GetStatusCode(), ctx.Path(), ctx.Method(), ctx.RemoteAddr())
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)

				logrus.Errorf("recover => %s", logMessage)

				ctx.StatusCode(500)
				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}

// 请求日志记录
func NewAccessLogMdw() iris.Handler {
	return func(ctx iris.Context) {
		begin := time.Now()

		reqBody := helper.RequestBody(ctx)
		// 如果请求内容不是json，则不记录
		if reqBody != "" && strings.Index(reqBody, "{") != 0 {
			reqBody = "non json body..."
		}

		ctx.Record()

		defer func() {
			respBody := string(ctx.Recorder().Body())
			// 响应内容必须是json格式，含有code码的字符串，否则忽略响应内容
			if strings.Index(respBody, "{") != 0 || strings.Index(respBody, "\"code\"") == -1 {
				respBody = "non json body..."
			}

			duration := time.Now().Sub(begin).Nanoseconds() / 1000000

			logrus.WithFields(logrus.Fields{
				"ip":     ctx.RemoteAddr(),
				"method": ctx.Method(),
				"path":   ctx.Path(),
				// 头信息的内容有写多，可以根据情况，只取某些字段
				// "header":   helper.RequestHeader(ctx),
				"queries":  helper.RequestQueries(ctx),
				"reqbody":  reqBody,
				"duration": duration,
			}).Info(respBody)
		}()

		ctx.Next()
	}
}