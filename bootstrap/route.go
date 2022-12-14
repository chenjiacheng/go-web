package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-web/app/http/middlewares"
	"go-web/pkg/response"
	"go-web/routes"
	"net/http"
	"strings"
)

func SetupRoute(r *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(r)

	// 注册 API 路由
	routes.RegisterAPIRoutes(r)

	// 配置 404 路由
	setup404Handler(r)
}

func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

// 处理 404 请求
func setup404Handler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			response.Abort404(c)
		}
	})
}
