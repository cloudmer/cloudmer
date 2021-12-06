package http

import "cloudmer/api/v1/controllers"

// 路由
func (service *httpService) route()  {
	// version 1
	v1 := service.ginEngine.Group("v1")
	{

		// demo
		demo := v1.Group("demo")
		{
			demo.GET("", controllers.DemoCreate)
		}

	}
}