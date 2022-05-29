package main

import (
	"daimao/log"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
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
	parse, _ := time.Parse("2006-01-02T15:04:05.000000000Z", "2022-05-27T03:48:42.553548987Z")
	fmt.Println(parse.Format("2006-01-02 15:04:05"))
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
