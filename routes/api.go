package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/app/controller"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		user := new(controller.User)
		v1.GET("/user/1", user.Get)
	}
}
