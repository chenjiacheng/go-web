package tests

import (
	"github.com/spf13/viper"
	"testing"
)

// 测试读取 env 配置文件
func TestEnv(t *testing.T) {

	v := viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig() // 搜索并读取配置文件
	if err != nil {
		panic(err)
	}

	t.Log("APP_PAGE_SIZE：", v.GetInt("APP_PAGE_SIZE"))
	t.Log("APP_APP_MODE：", v.GetString("APP_APP_MODE"))
}

// 测试读取 yaml 配置文件
func TestYaml(t *testing.T) {
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // 搜索并读取配置文件
	if err != nil {
		panic(err)
	}

	t.Log("PageSize：", viper.GetInt("app.PageSize"))
	t.Log("AppMode：", viper.GetString("app.AppMode"))
}
