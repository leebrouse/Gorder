package common

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/common/tracing"
	"github.com/sirupsen/logrus"
	"net/http"
)

// BaseResponse 是基础响应结构体，用于封装统一响应格式的逻辑方法（success / error）
type BaseResponse struct{}

// response 是统一响应格式的结构体
type response struct {
	Errno   int    `json:"errno"`    // 错误码，0 表示成功，非 0 表示失败
	Message string `json:"message"`  // 提示信息，成功为 "success"，失败为错误信息
	Data    any    `json:"data"`     // 响应数据内容，可以是任意类型
	TraceID string `json:"trace_id"` // 链路追踪 ID，用于定位请求问题
}

// Response 是对外暴露的统一响应函数，会根据是否出错调用 success 或 error
func (base *BaseResponse) Response(c *gin.Context, err error, data interface{}) {
	if err != nil {
		base.error(c, err) // 如果存在错误，走错误响应
	} else {
		base.success(c, data) // 否则走成功响应
	}
}

// success 封装成功响应的处理逻辑
func (base *BaseResponse) success(c *gin.Context, data interface{}) {
	traceID := tracing.TraceID(c.Request.Context())                        // 获取当前请求的链路追踪 ID
	logrus.WithField("trace_id", traceID).Info("handling success request") // 记录成功日志
	c.JSON(http.StatusOK, response{
		Errno:   0,
		Message: "success",
		Data:    data,
		TraceID: traceID,
	})
}

// error 封装错误响应的处理逻辑
func (base *BaseResponse) error(c *gin.Context, err error) {
	traceID := tracing.TraceID(c.Request.Context()) // 获取当前请求的链路追踪 ID
	c.JSON(http.StatusOK, response{
		Errno:   2,           // 错误码固定为 2（可根据项目需求自定义）
		Message: err.Error(), // 错误信息为 err 的内容
		Data:    nil,         // 错误时无返回数据
		TraceID: traceID,
	})
}
