package main

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/app"
)

func main() {

	config := app.InitConfig()
	fmt.Println(config)

	//an := console.Ask("Run build")
	//if an {
	//
	//	start := time.Now()
	//
	//	app.Create(config)
	//
	//	duration := time.Since(start)
	//
	//	console.Info("Build done")
	//	fmt.Println(duration)
	//
	//}
}