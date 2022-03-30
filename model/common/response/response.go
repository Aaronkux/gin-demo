package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	RELOGIN = 3
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, nil, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, nil, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, nil, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, nil, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
func ReloginWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(RELOGIN, data, message, c)
}

func FailWithCustomErrorOrDefault(message string, err error, c *gin.Context) {
	errMsg := message
	if cusError, ok := err.(*CusError); ok {
		errMsg = cusError.Error()
	}
	Result(ERROR, nil, errMsg, c)
}
