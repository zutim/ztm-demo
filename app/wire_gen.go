// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

// Injectors from wire.go:

func InitApps() *Apps {
	apps := NewApps()
	return apps
}
