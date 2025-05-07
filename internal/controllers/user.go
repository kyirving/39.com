package controllers

import (
	"fmt"
	"time"

	"39.com/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (c *UserController) Add(ctx *gin.Context) {
	resp := &utils.Response{Code: 200}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	resp.Output(ctx)
}
