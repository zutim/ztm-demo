package main

import (
	"github.com/zutim/ztm-demo/app"
	"sync"
)

func main() {
	app.InitApps()
	testDemo()
}

func testDemo() {
	res := 0

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		a := i
		wg.Add(1)
		app.App.Pool.Schedule(func() {
			defer wg.Done()
			res += a
		})
	}

	app.App.Pool.Stop()

	wg.Wait()

	app.App.Logs("").Info("game over！res is :", res)
}
