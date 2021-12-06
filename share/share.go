package share

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// console 日志输出
	Log *logrus.Logger
	// viper config
	Viper *viper.Viper
	// mysql
	Mysql *gorm.DB
)
