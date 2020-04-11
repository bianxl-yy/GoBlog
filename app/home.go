package app

import "github.com/bianxl-yy/GoBlog/app/handler"

func registerHomeHandler() {

	// 用户登录
	App.POST("/login", handler.Login)

	// 文章相关
	App.GET("/article/list", handler.Articles)
	App.GET("/article/list/:page/", handler.Articles)
	App.GET("/article/:slug", handler.Article)

	// 页面相关
	App.GET("/page/:slug", handler.Page)

	// 评论相关
	App.POST("/comment/:id/", handler.Comment)

	// 标签相关
	App.GET("/tag/:name/", handler.ArticlesByTag)
	App.GET("/tag/:name/:page/", handler.ArticlesByTag)

	// 获取配置
	App.GET("/config/get/:name", handler.GetConfig)
	App.GET("/config/get/all", handler.GetConfigAll)

	//App.GET("/feed", handler.Rss)
	//App.GET("/sitemap", handler.SiteMap)
}
