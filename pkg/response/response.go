package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"RequestId": Uniqid(""),
		"Code":      0,
		"Msg":       "Success",
		"Data":      data,
	})
}

func Uniqid(prefix string) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
}
