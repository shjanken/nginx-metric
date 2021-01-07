package main

import (
	"fmt"
	"log"
	"os"

	metric "com.github.shjanken/nginx-metric"
	"github.com/spf13/viper"
)

const (
	helpMessage = "Useage: ngxmetric [access.log]\n"

	testConfig = `
http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent"'

}
`
)

func main() {

	if len(os.Args) == 1 {
		displayUsage()
		os.Exit(1)
	}

	repo, err := initPostgresRepo()
	if err != nil {
		log.Fatalf("create postgres repo failure, %v", err)
	}

	if os.Args[1] == "initdb" {
		initDB(repo) // initDB 会重新创建数据库里面的表，如果有错误会中断程序的运行
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("open %s failure. %v\n", os.Args[1], err)
	}

	provider := metric.NewFileProvider(file, testConfig)
	service := metric.NewService(provider, repo)
	service.Save()
}

func displayUsage() {
	fmt.Println(helpMessage)
}

func initPostgresRepo() (postgre *metric.PostgresRepo, err error) {
	// read config file
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("ini")

	if err = viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file error %w", err)
	}

	return metric.NewPostgre(
		viper.GetString("database.dbname"),
		viper.GetString("database.user"),
		viper.GetString("database.password")), nil
}

func initDB(repo *metric.PostgresRepo) {
	// 如果命令行参数第一个是 initdb. 则在数据库里面重新创建表
	if err := repo.DropSchema(); err != nil {
		log.Fatal(err)
	}
	if err := repo.CreateSchema(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("create database schema success")
	os.Exit(0)
}
