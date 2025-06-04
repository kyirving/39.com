package routes

import (
	"39.com/internal/controllers"
	"39.com/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	// 初始化路由
	router := gin.Default()
	// 注册多个中间件
	router.Use(
		middleware.RequestCacheMiddleware(), // 请求缓存
		middleware.AccessLog(),              // 记录访问日志

		middleware.Middleware(), // 自定义中间件
	)

	RegisterUserRoutes(router)
	return router
}

func RegisterUserRoutes(r *gin.Engine) {
	user := controllers.NewUserController()
	r.POST("/user/add", user.Add)
}
