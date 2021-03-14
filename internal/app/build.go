package app

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/entities/php"
	"github.com/romaxa83/skeleton/internal/helpers"
)

var startDockerCompose string = `version: "3.7"
  services:
`

func Create(config *Config) {

	path := config.ProjectPath.Path + "/" + config.ProjectName.Name
	pathToDockerCompose := path + "/docker-compose.yml"

	// создаем папку проекта
	//helpers.CreateDir(path)

	// создаем файл docker-compose
	//helpers.AddToFile(pathToDockerCompose, startDockerCompose)

	// создаем сервер и прописываем docker-compose
	//serviceServer := server.Create(path, config.Server.Name, config.IP.IP, config.ProjectName.Name)
	//helpers.AddToFile(pathToDockerCompose, serviceServer)

	// создаем php и прописываем docker-compose
	servicePhp := php.Create(path, config.Php.Version, config.ProjectName.Name)
	helpers.AddToFile(pathToDockerCompose, servicePhp)

	fmt.Println(pathToDockerCompose)
}
