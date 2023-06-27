package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zutim/ztm-demo/network/ngin/handler"
)

func LoadRouter(r *gin.Engine) {
	hand := handler.NewHandler()
	r.POST("order/create", hand.OrderCreate)
}
