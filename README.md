## Usage
创建一个.ini结尾的文件，如config.ini，内容如下：
```$xslt
[server]
ip=127.0.0.1
port=90

[mysql]
username=root
passwd=123123
database=aa
host=127.0.0.1
port=3306
timeout=1.2
```

根据配置文件定义结构体：

```$xslt
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
```

通过UnmarshalFile方法即可将文件内容加载到定义的Config结构体中
```$xslt
config := &Config{}
err := iniConfig.UnmarshalFile(fileName, config)
```

同样的也可以通过MarshalFile方法将结构体的内容写入到文件中
```$xslt
fileData, err := ioutil.ReadFile("./config.ini")
if err != nil {
    t.Error("open file error:",err)
}
var config Config
err = unmarshal(fileData, &config)
```



