package main

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/app"
	"github.com/romaxa83/skeleton/internal/console"
	"time"
)

func main() {

	console.Info("CREATOR APP")
	//выбираем язык
	local := app.ChoiceLocal()

	config := app.InitConfig(local)

	//fmt.Println(config.Server)

	an := console.Ask("Run build")
	if an {

		start := time.Now()

		// запускается сборка проекта
		app.Create(config)

		duration := time.Since(start)

		console.Info("Build done")
		fmt.Println(duration)

	}
}