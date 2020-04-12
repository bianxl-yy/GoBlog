package app

import (
	"fmt"
	"github.com/bianxl-yy/GoBlog/app/model"
	"github.com/bianxl-yy/GoBlog/app/utils"
	"github.com/labstack/echo"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

var (
	// APP VERSION, as date version
	VERSION = 20140228
	// Global GoInk application
	//App *GoInk.App
	App *echo.Echo
)

//var staticFileSuffix = ".css,.js,.jpg,.jpeg,.png,.gif,.ico,.xml,.zip,.txt,.html,.otf,.svg,.eot,.woff,.ttf,.doc,.ppt,.xls,.docx,.pptx,.xlsx,.xsl"

func init() {
	// init application
	App = echo.New()

	//App.Use(middleware.CORS())

	// TODO: 暂未限制文件类型
	App.Static("/static", model.Config.String("static_dir"))

	// set recover handler
	/*
		App.Recover(func(context *GoInk.Context) {
			go LogError(append(append(context.Body, []byte("\n")...), debug.Stack()...))
			theme := handler.Theme(context)
			if theme.Has("error/error.html") {
				theme.Layout("").Render("error/error", map[string]interface{}{
					"error":   string(context.Body),
					"stack":   string(debug.Stack()),
					"context": context,
				})
			} else {
				context.Body = append([]byte("<pre>"), context.Body...)
				context.Body = append(context.Body, []byte("\n")...)
				context.Body = append(context.Body, debug.Stack()...)
				context.Body = append(context.Body, []byte("</pre>")...)
			}
			context.End()
		})

		// set not found handler
		App.NotFound(func(context *GoInk.Context) {
			theme := handler.Theme(context)
			if theme.Has("error/notfound.html") {
				theme.Layout("").Render("error/notfound", map[string]interface{}{
					"context": context,
				})
			}
			context.End()
		})
	*/

	// add recover defer
	defer func() {
		e := recover()
		if e != nil {
			bytes := append([]byte(fmt.Sprint(e)+"\n"), debug.Stack()...)
			utils.LogError(bytes)
			println("panic error, crash down")
			os.Exit(1)
		}
	}()

	// catch exit command
	go catchExit()
}

// code from https://github.com/Unknwon/gowalker/blob/master/gowalker.go
func catchExit() {
	sigTerm := syscall.Signal(15)
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, sigTerm)

	for {
		switch <-sig {
		case os.Interrupt, sigTerm:
			// TODO: 待添加，将配置写入文件
			println("ready to exit")
			os.Exit(0)
		}
	}
}

// Init starts Fxh.Go application preparation.
// Load models and plugins, update views.
func Init() {
	/*
			// TODO: 待处理
			// init plugin
			plugin.Init()

			// TODO: 待处理
			// update plugin handlers
			plugin.Update(App)

		// TODO: 待处理
		println("app version @ " + strconv.Itoa(model.GetVersion().Version))
	*/
}

// Run begins Fxh.Go http server.
func Run() {
	// 注册后台路由
	registerAdminHandler()
	// TODO: 待处理
	//registerCmdHandler()
	// 注册前端路由
	registerHomeHandler()
	// 启动服务
	App.Logger.Fatal(App.Start(":9001"))
}
