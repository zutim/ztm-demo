//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
)

func InitApps() *apps {
	panic(wire.Build(newApps))
}
