package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zutim/http"
	"github.com/zutim/ztm-demo/common/apis"
	"github.com/zutim/ztm-demo/common/pb/order"
	"github.com/zutim/ztm-demo/common/pb/user"
	"github.com/zutim/ztm-demo/network/ngin/handler"
	"github.com/zutim/ztm-demo/network/ngin/middleWare"
	"google.golang.org/grpc"
)

func GetRpcRouter() func(server *grpc.Server) {
	return func(server *grpc.Server) {
		order.RegisterOrderServer(server, apis.Order())
		user.RegisterUserServer(server, apis.User())
	}
}

func GetRouter() func(r *gin.Engine) {
	hand := handler.NewHandler()
	groups := []http.RouteGroup{
		{
			Prefix: "v1",
			Routes: []http.Route{
				{Method: "POST", URL: "order/create", Handler: hand.OrderCreate},
				{Method: "POST", URL: "order/create3", Handler: hand.OrderCreate3},
				{Method: "POST", URL: "order/create4", Handler: hand.OrderCreate4},
			},
			Handlers: []gin.HandlerFunc{middleWare.Cors()},
		},
		{
			Prefix: "v2",
			Routes: []http.Route{
				{Method: "POST", URL: "order/create2", Handler: hand.OrderCreate2},
			},
		},
	}

	return func(r *gin.Engine) {
		for _, group := range groups {
			http.RegisterRouteGroup(r, group)
		}
	}
}

func GetRouter2() func(r *gin.Engine) {
	hand := handler.NewHandler()
	return func(r *gin.Engine) {
		r2 := r.Group("test").Use(middleWare.Cors())
		r2.POST("order/create", hand.OrderCreate)
	}
}
