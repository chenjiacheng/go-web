package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/app/controller"
	"go-web/app/middleware"
	"io"
	"net/http"
	"os"
)

/*type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}*/

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		user := new(controller.User)
		v1.GET("/user/1", user.Get)
	}

	/*authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})*/

	r.Use(middleware.Sign())

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.GET("/ping", func(c *gin.Context) {
		/*c.JSON(200, gin.H{
			"message": "pong",
		})*/

		user := c.Query("user")
		fmt.Println("user=" + user)

		/*var form LoginForm
		c.ShouldBind(&form)

		fmt.Println("User=" + form.User)
		fmt.Println("Password=" + form.Password)*/

		/*data := map[string]interface{}{
			"lang": "go语言",
			"tag":  1,
		}*/

		// data := []string{"lena", "austin", "foo"}

		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		// c.JSON(200, data)
		// c.AsciiJSON(http.StatusOK, data)
		// c.JSONP(http.StatusOK, data)
		c.SecureJSON(http.StatusOK, msg)

	})
	r.Run()
}
