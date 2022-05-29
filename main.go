package main

import (
	"daimao/log"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logger.Info("自定义logger", zap.String("name", "zap log"))
	logger.Debug("自定义logger", zap.String("name", "zap log"))
	fmt.Println(viper.AllSettings())
	fmt.Println(viper.GetString("pylogon.path"))
	fmt.Println(viper.GetStringSlice("pylogon.symbols"))
	fmt.Println(viper.GetInt("pylogon.limit"))
	//fmt.Print("请输入")
	//var cmd string
	//fmt.Scan(&cmd)
	//fmt.Println(cmd)
	//time.Sleep(1 * time.Second)
}

func init() {
	loadConfig()
	logger = log.GetLogger()
}

func loadConfig() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
