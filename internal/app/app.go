package app

import (
	"fmt"
	"github.com/labstack/echo"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/bianxl-yy/GoBlog/internal/app/article/content"
	"github.com/bianxl-yy/GoBlog/internal/app/article/tag"
	"github.com/bianxl-yy/GoBlog/internal/app/comment"
	"github.com/bianxl-yy/GoBlog/internal/app/page"
	setting "github.com/bianxl-yy/GoBlog/internal/app/system/setting/base"
	"github.com/bianxl-yy/GoBlog/internal/app/user/auth"
	"github.com/bianxl-yy/GoBlog/internal/pkg/storage"
	"github.com/bianxl-yy/GoBlog/internal/pkg/utils"
)

var (
	// Global GoInk application
	App *echo.Echo
)

func init() {
	// init application
	App = echo.New()

	// 初始化设置
	setting.Init()
	// 初始化用于保存数据的... 工具...
	storage.Init(setting.Config.String("storage"))

	//App.Use(middleware.CORS())

	// TODO: 暂未限制文件类型
	App.Static("/static", setting.Config.String("static_dir"))

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


func registerHandler() {

	// 用户登录
	App.POST("/login", auth.Login)

	// 文章相关
	App.GET("/article/list", content.Articles)
	App.GET("/article/list/:page/", content.Articles)
	App.GET("/article/:slug", content.Get)

	// 页面相关
	App.GET("/page/:slug", page.Get)

	// 评论相关
	App.POST("/comment/:id/", comment.Get)

	// 标签相关
	App.GET("/tag/:name/", tag.Articles)
	App.GET("/tag/:name/:page/", tag.Articles)

	// 获取配置
	App.GET("/setting/get/:name", base.Get)
	App.GET("/setting/get", base.GetAll)

}

// Run begins Fxh.Go http server.
func Run() {
	// 注册前端路由
	registerHandler()
	// 启动服务
	App.Logger.Fatal(App.Start(":9001"))
}
