package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// set default config value
	viper.SetDefault("default", "defalut_value")

	// read config file
	viper.SetConfigName("metric")
	viper.SetConfigType("ini")

	viper.AddConfigPath("$HOME/metric/")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failure", err)
	}

	fmt.Println("port is: ", viper.GetInt("database.port"))
	fmt.Println("host is: ", viper.GetString("database.host"))
	fmt.Println("default is: ", viper.GetString("default"))
}
