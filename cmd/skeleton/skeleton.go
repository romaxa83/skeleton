package main

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/app"
	"github.com/romaxa83/skeleton/internal/app/build"
	"github.com/romaxa83/skeleton/internal/app/config"
	"github.com/romaxa83/skeleton/internal/console"
	"os"
	"time"
)

func main() {

	console.Info("CREATOR APP")
	// выбираем стратегию для приложения
	strategy := app.ChoiceStrategy()

	if app.IsMonolit(strategy) {
		console.Info("Create monolit")
		//config := config.InitMonolitConfig()
		_ = config.InitMonolitConfig()


	} else if app.IsMicroservice(strategy) {
		console.Info("Create microservice")

		config := config.InitMicroserviceConfig()
		an := console.Ask("Run build")
		if an {

			start := time.Now()
			// запускается сборка проекта
			build.CreateMicroservice(config)

			duration := time.Since(start)

			console.Info("Build done")
			fmt.Println(duration)

		}
	} else {
		console.Info("Create one service (developing)")

		os.Exit(2)
	}
	
	fmt.Println(strategy.Type)
	//os.Exit(2)

	//local := app.ChoiceLocal()

	//config := app.InitConfig(local)


	//an := console.Ask("Run build")
	//if an {
	//
	//	start := time.Now()
	//
	//	// запускается сборка проекта
	//	app.Create(config)
	//
	//	duration := time.Since(start)
	//
	//	console.Info("Build done")
	//	fmt.Println(duration)
	//
	//}
}