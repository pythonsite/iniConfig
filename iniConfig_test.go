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
	fileData, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("open file error:",err)
	}
	var config Config
	err = unmarshal(fileData, &config)
	if err != nil {
		t.Fatalf("unmarshal failed, err:%v",err)
	}
	t.Logf("unmarshal success config:%#v", config)

}

// 测试如果传入的不是config不是指针的时候的异常
func TestUnmarshalConfigNOTPtr(t *testing.T) {
	fileData, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("open file error:",err)
	}
	var config Config
	err = unmarshal(fileData, &config)
	if err != nil {
		t.Fatalf("unmarshal failed,err:%v",err)
	}
	t.Logf("unmarshal success config:%#v", config)
}


func TestUnmarshalConfig(t *testing.T) {
	fileData, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("open file error:",err)
	}
	var config Config
	err = unmarshal(fileData, &config)
	if err != nil {
		t.Fatalf("unmarshal failed,err:%v",err)
	}
	t.Logf("unmarshal success config:%#v", config)

}

func TestMarshalConfig(t *testing.T) {
	fileData, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("open file error:",err)
	}
	var config Config
	err = unmarshal(fileData, &config)
	if err != nil {
		t.Fatalf("unmarshal failed,err:%v",err)
	}
	t.Logf("unmarshal success config:%#v", config)

	res , err := marshal(config)
	if err != nil {
		t.Fatalf("marshal error:%v", err)

	}
	t.Logf("marshal success res :%#v", res)
}

// 测试将结构体内容序列化到文件中
func TestMarshalFile(t *testing.T) {
	fileData, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("open file error:",err)
	}
	var config Config
	err = unmarshal(fileData, &config)
	if err != nil {
		t.Fatalf("unmarshal failed,err:%v",err)
	}
	t.Logf("unmarshal success config:%#v", config)
	err = MarshalFile("./test.ini", config)
	if err != nil {
		t.Fatalf("MarshalFile error:%v", err)
	}
	t.Logf("MarshalFile success")
}

func TestUnmarshalFile(t *testing.T) {
	var config Config
	err := UnmarshalFile("./config.ini", &config)
	if err != nil {
		t.Fatalf("UnmarshalFile failed err:%v",err)
	}
	t.Logf("UnmarshalFile success,Config:%#v",config)
}



