package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	conf "github.com/zutim/config"
	"github.com/zutim/ztm-demo/app"
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

	testDemo()
}

func testDemo() {

	type Ticket struct {
		Id     string `json:"id"`
		AutoId int    `json:"auto_id"`
	}

	var t Ticket
	if err := app.DB().Table("tbl_tickets").First(&t).Error; err != nil {
		app.Log().Error(err)
		return
	}

	fmt.Println(t)

	//res := 0
	//
	//wg := sync.WaitGroup{}
	//
	//for i := 0; i < 1000; i++ {
	//	a := i
	//	wg.Add(1)
	//	app.App.Pool.Schedule(func() {
	//		defer wg.Done()
	//		res += a
	//	})
	//}
	//
	//app.App.Pool.Stop()
	//
	//wg.Wait()
	//
	//app.App.Logs("").Info("game over！res is :", res)
}
