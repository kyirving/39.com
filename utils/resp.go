package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS            = 200 // 成功
	PARAMS_ERR         = 400 // 参数错误
	UNAUTHORIZED       = 401 // 未授权
	NOT_LOGIN          = 402 // 未登录
	FORBIDDEN          = 403 // 禁止访问
	NOT_FOUND          = 404 // 未找到
	METHOD_NOT_ALLOWED = 405 // 方法不允许
	REQUEST_TIMEOUT    = 408 // 请求超时
	TOO_MANY_REQUESTS  = 429 // 太多请求
	SERVER_FAIL        = 500 // 服务器内部错误
	BAD_GATEWAY        = 502 // 错误的网关
	GATEWAY_TIMEOUT    = 504 // 网关超时
)

var msgMaps = map[int]string{
	SUCCESS:            "success",
	SERVER_FAIL:        "服务器异常",
	NOT_FOUND:          "请求资源不存在",
	FORBIDDEN:          "请求资源被禁止",
	PARAMS_ERR:         "参数错误",
	UNAUTHORIZED:       "未授权",
	NOT_LOGIN:          "未登录",
	METHOD_NOT_ALLOWED: "请求方法不允许",
	REQUEST_TIMEOUT:    "请求超时",
	TOO_MANY_REQUESTS:  "请求过于频繁",
	BAD_GATEWAY:        "上下游服务异常",
	GATEWAY_TIMEOUT:    "上下游服务超时",
}

func (r *Response) Output(ctx *gin.Context) {

	if r.Msg == "" {
		if value, ok := msgMaps[r.Code]; ok {
			r.Msg = value
		}
	}
	ctx.JSON(SUCCESS, r)
	ctx.Abort() // 阻止后续的处理函数执行
	return
}
