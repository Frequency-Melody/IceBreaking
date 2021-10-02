package config

import (
	_ "embed"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

var c *config

func Get() *config {
	return c
}

func init() {
	// 初始化配置
	initConfig()
	// 检验配置参数合法性
	// 打算废弃，因为，可以直接转 Json 判断。。。。
	// 目前实现的配置检查只支持单级，多级直接用 json 、map 多香啊
	if !Get().validConfig(validRule) {
		fmt.Println("配置文件有误")
		os.Exit(1)
	}
}

//go:embed config.toml
var configFilename string

func initConfig() {
	c = new(config)
	if _, err := toml.Decode(configFilename, &c); err != nil {
		panic(err)
	}
	// 从环境变量读取敏感信息
	c.Mysql.Host = os.Getenv("MYSQL_HOST")
	c.Mysql.Pwd = os.Getenv("MYSQL_PWD")

	//c.OSS.EndPoint = os.Getenv("OSS_ENDPOINT")
	c.OSS.AccessKeyId = os.Getenv("OSS_ACCESS_KEY_ID")
	c.OSS.AccessKeySecret = os.Getenv("OSS_ACCESS_KEY_SECRET")

	// 杭电助手鉴权相关
	c.Hduhelp.ClientId = os.Getenv("HDUHELP_CLIENT_ID")
	c.Hduhelp.ClientSecret = os.Getenv("HDUHELP_CLIENT_SECRET")

}
