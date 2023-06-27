package order

import (
	"github.com/google/wire"
	"github.com/zutim/ztm-demo/internal/order/data"
	"github.com/zutim/ztm-demo/internal/order/service"
)

var ProviderSet = wire.NewSet(data.NewOrderData, service.NewOrderService)
