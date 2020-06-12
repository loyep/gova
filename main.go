package main

import (
	"gova/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// 跨域
	// Router.Use(middleware.Cors())
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitApiRouter(ApiGroup) // 注册功能api路由
	return Router
}

func main() {

}
