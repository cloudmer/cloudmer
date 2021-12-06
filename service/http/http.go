package http

import (
	"cloudmer/share"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type httpService struct {
	// 端口
	Port string
	// gin engine
	ginEngine *gin.Engine
}

// http run
func (service *httpService) Build()  {
	// route 路由设置
	service.route()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	service.ginEngine.Run(service.Port)
}

func StdConfig(key string) *httpService {
	httpService := DefaultConfig()
	if err := share.Viper.UnmarshalKey(key, httpService); err != nil {
		logrus.Fatal(err)
	}
	return httpService
}

// 默认配置项
func DefaultConfig() *httpService {
	return &httpService{
		//Port: "0.0.0.0:8080",
		Port: "8080",
		ginEngine: gin.Default(),
	}
}
