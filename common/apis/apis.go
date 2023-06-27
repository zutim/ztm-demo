package apis

import (
	"github.com/zutim/ztm-demo/common/pb/order"
	"github.com/zutim/ztm-demo/common/pb/user"
)

type Api struct {
	user.UserServer
	order.OrderServer
}

var (
	apis *Api
)

func NewApis(user user.UserServer, order order.OrderServer) *Api {
	apis = &Api{
		user,
		order,
	}
	return apis
}

func User() user.UserServer {
	return apis.UserServer
}

func Order() order.OrderServer {
	return apis.OrderServer
}
