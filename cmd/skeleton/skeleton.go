package main

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/app"
	"github.com/romaxa83/skeleton/internal/console"
	"time"
)

func main() {

	console.Info("CREATOR APP")
	config := app.InitConfig()

	an := console.Ask("Run build")
	if an {

		start := time.Now()

		app.Create(config)

		duration := time.Since(start)

		console.Info("Build done")
		fmt.Println(duration)

	}
}