package main

import (
	"fmt"
    "github.com/spf13/viper"
)

// 结构体变量的名字必须和yaml配置里的名字一致，否则加标签也没用
type Config struct {
	ListenOn SvcConf  `yaml:"ListenOn"`
	Redis RedisConf  
	Mysql MysqlConf 
}

type SvcConf struct {
	Host string
	Port int
}

type RedisConf struct {
	Addrs []string
}

type MysqlConf struct {
	Host string
	Port int
	Database string
	User string
	Password string
}

func main() {
	var config Config
	// viper.SetConfigFile("./config.yaml")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
	viper.SetDefault("ListenOn.Host", "localhost")
	viper.SetDefault("ListenOn.Port", 8080)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	viper.Unmarshal(&config)
	viper.Set("Mysql.Host", "127.0.0.1")
	viper.WriteConfig()
	fmt.Println(config)
}