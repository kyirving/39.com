package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SuccessCode = 200
	ErrorCode   = 500
)

var msgMaps = map[int]string{
	SuccessCode: "success",
	ErrorCode:   "error",
}

func (r *Response) Output(ctx *gin.Context) {

	if r.Msg == "" {
		if value, ok := msgMaps[r.Code]; ok {
			r.Msg = value
		}
	}
	ctx.JSON(SuccessCode, r)
	ctx.Abort() // 阻止后续的处理函数执行
}
