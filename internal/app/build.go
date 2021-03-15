package app

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/entities/db"
	"github.com/romaxa83/skeleton/internal/entities/php"
	"github.com/romaxa83/skeleton/internal/entities/server"
	"github.com/romaxa83/skeleton/internal/helpers"
	"strings"
)

var startDockerCompose string = `version: "3.7"
  services:
`

var endDockerCompose string = `networks:
  default:
    driver: bridge
    ipam:
      config:
        - subnet: {ip}
`

func Create(config *Config) {

	path := config.ProjectPath.Path + "/" + config.ProjectName.Name
	//pathToDockerCompose := path + "/docker-compose.yml"

	// создаем папку проекта
	helpers.CreateDir(path)

	// создаем docker-compose с выбраными сервисами, а также создаем под них конфиги
	createDockerCompose(config, path)


	//fmt.Println(pathToDockerCompose)
}

// с
func createDockerCompose(config *Config, path string)  {
	pathToDockerCompose := path + "/docker-compose.yml"
	fmt.Println(pathToDockerCompose)

	// создаем файл docker-compose
	helpers.AddToFile(pathToDockerCompose, startDockerCompose)

	// создаем сервер и прописываем docker-compose
	serviceServer := server.Create(path, config.Server.Name, config.IP.IP, config.ProjectName.Name)
	helpers.AddToFile(pathToDockerCompose, serviceServer)

	// создаем php и прописываем docker-compose
	servicePhp := php.Create(path, config.Php.Version, config.ProjectName.Name)
	helpers.AddToFile(pathToDockerCompose, servicePhp)

	// создаем бд и прописываем docker-compose
	serviceDb := db.Create(
		path,
		config.DB.Driver,
		config.IP.IP,
		config.ProjectName.Name,
		config.DB.Version,
		config.DB.UserName,
		config.DB.Password,
		config.DB.DbName,
	)
	helpers.AddToFile(pathToDockerCompose, serviceDb)

	// здесь docker-compose создан
	helpers.AddToFile(pathToDockerCompose, strings.Replace(endDockerCompose, "{ip}", helpers.NetworkIP(config.IP.IP),1))
}
