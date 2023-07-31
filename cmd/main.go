package main

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	conf "github.com/zutim/config"
	"github.com/zutim/http"
	ztm_demo "github.com/zutim/ztm-demo"
	"github.com/zutim/ztm-demo/common/app"
	"github.com/zutim/ztm-demo/network/ngin/router"
)

var (
	config []string
	arg    string
)

func init() {
	flag.StringArrayVar(&config, "configPath", []string{"config/db.yaml"}, "default config path.")
	flag.StringVar(&arg, "argStatus", "start", "os.arg .")
}

func main() {

	conf.LoadConfig(config)

	app.InitApps()

	ztm_demo.InitApis()

	run()
}

func run() {
	//
	//go grpc.NewGrpc().Register(router.GetRpcRouter()).Run(viper.GetString("rpc"))
	//
	//http.NewHttp().Register(router.GetRouter()).Run(viper.GetString("http"))
	http.NewHttp().Register(router.GetRouter2()).Run(viper.GetString("http"))
}
