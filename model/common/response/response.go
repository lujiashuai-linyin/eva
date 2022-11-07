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
	ERROR   = 1
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
func CResult(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusBadRequest, Response{
		code,
		data,
		msg,
	})
}
func SResult(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusInternalServerError, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	CResult(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	CResult(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	CResult(ERROR, data, message, c)
}

func SFail(c *gin.Context) {
	CResult(ERROR, map[string]interface{}{}, "操作失败", c)
}

func SFailWithMessage(message string, c *gin.Context) {
	CResult(ERROR, map[string]interface{}{}, message, c)
}

func SFailWithDetailed(data interface{}, message string, c *gin.Context) {
	CResult(ERROR, data, message, c)
}
