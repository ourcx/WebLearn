package config

//config就是定义了一个基础的包，这个包由很多配置
import "github.com/spf13/viper"

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	// 配置文件类型名字
	viper.SetConfigType("yaml")
	//设置配置文件名字是yaml
	viper.AddConfigPath("./config")
	//添加配置文件搜索路径（当前目录下的config子目录）

	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	//初始化全局变量
	AppConfig = &Config{}
	// 绑定到结构体
	if err := viper.Unmarshal(AppConfig); err != nil {
		panic(err)
	}

	initDB()
	initRedis()
}
