package routes

import (
	"39.com/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	// 初始化路由
	router := gin.Default()
	RegisterUserRoutes(router)
	return router
}

func RegisterUserRoutes(r *gin.Engine) {
	user := &controllers.UserController{}
	r.POST("/user/add", user.Add)
}
