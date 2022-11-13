package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"request_id": uniqid(""),
		"code":       0,
		"msg":        "success",
		"data":       data,
	})
}

func Abort404(c *gin.Context, msg ...string) {
	c.JSON(http.StatusOK, gin.H{
		"request_id": uniqid(""),
		"code":       404,
		"msg":        defaultMessage("路由未定义，请确认 url 和请求方法是否正确。", msg...),
	})
}

func uniqid(prefix string) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
}

func defaultMessage(defaultMsg string, msg ...string) string {
	message := defaultMsg
	if len(msg) > 0 {
		message = msg[0]
	}
	return message
}
