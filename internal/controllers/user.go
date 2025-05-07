package controllers

import (
	"39.com/internal/dto"
	"39.com/internal/model"
	"39.com/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	*BaseController
}

// NewUserController 构造函数
func NewUserController() *UserController {
	return &UserController{
		BaseController: NewBaseController(),
	}
}

func (c *UserController) Add(ctx *gin.Context) {
	RegisterBody := &dto.Register{}
	resp := &utils.Response{}
	if err := ctx.ShouldBind(RegisterBody); err != nil {
		c.logger.Error(err)
		resp.Code = utils.PARAMS_ERR
		resp.Output(ctx)
		return
	}

	// 业务逻辑
	user := model.NewUserModel()
	user.Username = RegisterBody.Username
	user.Password = RegisterBody.Password
	err := user.Add()
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"username": user.Username,
			"password": user.Password,
		}).Error(err)
		resp.Code = utils.SERVER_FAIL
		resp.Output(ctx)
		return
	}

	resp.Code = utils.SUCCESS
	resp.Output(ctx)
}
