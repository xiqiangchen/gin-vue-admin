package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
	Illegal = -1
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
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		nil,
		message,
	})
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func IllegalWithMessage(message string, c *gin.Context) {
	Result(Illegal, map[string]interface{}{}, message, c)
}

func OkWithNoContent(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func NoContent(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}

func AutoContent(resp any, c *gin.Context) {
	// 暂不处理压缩
	contentType := strings.ToLower(c.Request.Header.Get("content-type"))
	switch {
	case strings.Contains(contentType, "application/json"):
		c.JSON(http.StatusOK, resp)
	default:
		c.ProtoBuf(http.StatusOK, resp)
	}

}
