package response

import "github.com/gin-gonic/gin"

// 成功
func Success(msg string, data interface{}) gin.H {
	return gin.H{
		"code": CODE_SUCCESS,
		"msg": 	msg,
		"data": data,
	}
}

// 错误
func Error(err error) gin.H {
	return gin.H{
		"code": CODE_ERROR,
		"msg": 	err.Error(),
		"data": nil,
	}
}

// 聪明的适配
func Clever(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg": 	msg,
		"data": data,
	}
}