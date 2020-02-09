package main

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/middleware/logger"
	//"github.com/kataras/iris/middleware/recover"
	"iris_test/router"
	"iris_test/config"
)

func main() {
	app := iris.New()
	//app.Use(recover.New())
	//app.Use(logger.New())

	//设置视图目录
	app.RegisterView(iris.HTML("./views", ".html"))
	// 设置静态资源目录
	app.HandleDir("/static", "static")
	//初始化路由
	router.InitRouter(app)

	//从配置文件中读取端口，运行服务
	app.Run(iris.Addr(":" + config.GetString("httpport", "8080")))
}