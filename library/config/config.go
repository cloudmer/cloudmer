package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	instance *viper.Viper
)

var once sync.Once

func GetInstance() *viper.Viper {
	once.Do(func() {
		instance = viper.New()
	})
	return instance
}

func Loader(path string)  {
	viper := GetInstance()
	// 设置配置文件路径
	viper.SetConfigFile(path)
	// 读取该配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
