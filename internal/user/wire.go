package user

import (
	"github.com/google/wire"
	"github.com/zutim/ztm-demo/internal/user/service"
)

var ProviderSet = wire.NewSet(service.NewUserService)
