package main

import(
    "fmt"
    "github.com/spf13/viper"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

type User struct {
	Username string `gorm:"column:name"`
	Password string
}

// 读取配置文件config
type Config struct {
    Redis string
    MySQL MySQLConfig
}

type MySQLConfig struct {
    Port int
    Host string
    Username string
    Password string
	Database string
}

func main() {
    // 把配置文件读取到结构体上
    var config Config
    
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println(err)
        return
    }
     
    viper.Unmarshal(&config) //将配置文件绑定到config上
    
	mysqlDNS := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port, config.MySQL.Database)
	fmt.Println(mysqlDNS)
	// db, err := gorm.Open(mysql.Open(mysqlDNS))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var user = User{
	// 	Username : "test",
	// 	Password : "test",
	// }

	// db.Create(&user)
}