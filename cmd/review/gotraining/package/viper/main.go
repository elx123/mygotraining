package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type MySQLConfig struct {
	URL      string
	Username string
	Password string
}

type Config struct {
	Port  int
	MySQL MySQLConfig
}

func main() {
	var config Config
	viper.SetConfigName("config2") // 设置配置文件名 (不带后缀)
	//viper.SetConfigName("config2") // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(".")    // 第一个搜索路径
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&config) // 将配置信息绑定到结构体上
	fmt.Println(config)
}
