package main

import (
	flag "github.com/spf13/pflag"
	conf "github.com/zutim/config"
	ztm_demo "github.com/zutim/ztm-demo"
	"github.com/zutim/ztm-demo/common/app"
	"github.com/zutim/ztm-demo/network/server"
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

	//go server.NewGrpc().Register().Run()

	server.NewHttp().Run()

}
