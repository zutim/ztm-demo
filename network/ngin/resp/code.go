package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/zutim/ztm-demo/network/ngin/respCode"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": respCode.SUCCESS,
		"msg":  "success",
		"data": data,
	})
}

func FailResponse(c *gin.Context, msg string, code int) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}

func FailResponse1(c *gin.Context, msg string, code int, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
