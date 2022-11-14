package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/bootstrap"
	btsConfig "go-web/config"
	"go-web/pkg/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	r := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
