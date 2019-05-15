package iniConfig

import (
	"io/ioutil"
	"testing"
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

func TestUnmarshal(t *testing.T) {
	fileData, err := ioutil.ReadFile("./confi.ini")
	if err != nil {
		t.Error("open file error:",err)
	}
	var config Config
	err = unmarshal(fileData, config)
	if err != nil {
		t.Fatalf("unmarshal failed, err:%v",err)
	}

}

