//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
)

func InitApps() *Apps {
	panic(wire.Build(NewApps))
}
