package controllers

import (
	"39.com/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (c *UserController) Add(ctx *gin.Context) {
	resp := &utils.Response{Code: 200}
	resp.Output(ctx)
}
