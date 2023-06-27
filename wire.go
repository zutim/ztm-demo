//go:build wireinject
// +build wireinject

package ztm_demo

import (
	"github.com/google/wire"
	"github.com/zutim/ztm-demo/common/apis"
	"github.com/zutim/ztm-demo/internal/order"
	"github.com/zutim/ztm-demo/internal/user"
)

func InitApis() *apis.Api {
	panic(wire.Build(order.ProviderSet, user.ProviderSet, apis.NewApis))
}
