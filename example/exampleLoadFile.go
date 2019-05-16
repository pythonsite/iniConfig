package main


import (
	"fmt"
	"iniConfig"
)

type ServerConfig struct {
	Ip string	`ini:"ip"`
	Port int	`ini:"port"`
}

type MysqlConfig struct {
	UserName string	`ini:"username"`
	Passwd string	`ini:"passwd"`
	DataBase string `ini:"database"`
	Host string	`ini:"host"`
	Port int	`ini:"port"`
	Timeout float32 `ini:"timeout"`
}

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf MysqlConfig	`ini:"mysql"`
}

func loadFile(fileName string) {
	config := &Config{}
	err := iniConfig.UnmarshalFile(fileName, config)
	if err != nil {
		fmt.Println("UnmarshalFile error:",err)
		return
	}
	fmt.Printf("UnmarshalFile success:%#v", config)
}

func main() {
	loadFile("./test.ini")
}