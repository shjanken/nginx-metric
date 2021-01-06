package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// set default config value
	viper.SetDefault("default", "defalut_value") // 设置默认值

	// read config file
	// 配置文件的文件名和类型
	// viper 可以通过文件名和类型自动解析文件
	viper.SetConfigName("metric")
	viper.SetConfigType("ini")

	// 可以设置多个配置文件的路径
	// 后面路径中的配置项目会覆盖前面的
	viper.AddConfigPath("$HOME/metric/")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failure", err)
	}

	// 可以通过 . 符号获取配置项目的内容
	fmt.Println("port is: ", viper.GetInt("database.port"))
	fmt.Println("host is: ", viper.GetString("database.host"))
	fmt.Println("default is: ", viper.GetString("default"))

	// vipper 可以读取 env
	// vipper 读取环境变量的时候会自动大写
	viper.BindEnv("foo") // 将变量绑定到 key 上

	os.Setenv("FOO", "foo value") // 设置系统环境变量必须是大写
	foo := viper.Get("foo")       // 读取时是小写
	fmt.Println("foo environment is: ", foo)
}
