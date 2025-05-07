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
	router.Use(middleware.AccessLog(), middleware.Middleware())

	RegisterUserRoutes(router)
	return router
}

func RegisterUserRoutes(r *gin.Engine) {
	user := controllers.NewUserController()
	r.POST("/user/add", user.Add)
}
