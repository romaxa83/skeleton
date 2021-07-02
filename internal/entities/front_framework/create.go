package front_framework

import (
	"github.com/romaxa83/skeleton/internal/console"
	"github.com/romaxa83/skeleton/internal/entities/front_framework/inner_helpers"
	"github.com/romaxa83/skeleton/internal/helpers"
	"strings"
)

var localDockerCompose string = `version: "3.7"
services:
  front:
    build:
      context: docker
      dockerfile: dev/node/Dockerfile
    volumes:
      - ./:/app
    tty: true

networks:
  default:
    driver: bridge
    ipam:
      config:
        - subnet: 192.168.180.0/24
`
var dockerCompose string = `
  {name}-nginx:
    build:
      context: ./{name}/docker
      dockerfile: dev/nginx/Dockerfile
    container_name: {pname}__{name}_nginx
    hostname: {pname}__{name}_nginx
    volumes:
      - ./{name}:/app
    ports:
      - {ip}:8080:80

  {name}-node:
    build:
      context: ./{name}/docker
      dockerfile: dev/node/Dockerfile
    container_name: {pname}__{name}_node
    hostname: {pname}__{name}_node
    volumes:
      - ./{name}:/app
    working_dir: /app
    tty: true

`

func Create(
	data *Data,
	path string,
	ip string,
	projectName string,
) string {
	// название папки, где будет создан проект
	tempName := "test"

	// переменые путей
	pathToFrontService := path + "/" + data.Name
	pathToFrontServiceDockerDEV := pathToFrontService + "/docker/dev"
	pathToLocalDockerComposeFile := pathToFrontService + "/docker-compose.yml"


	// создаем папки и файлы для конфигов
	helpers.CreateDir(pathToFrontService)
	//helpers.CreateDir(pathToFrontServiceDocker)

	// создаем внутрений docker-compose, чтоб поднять node , и запустить установку фронта
	helpers.AddToFile(pathToLocalDockerComposeFile, localDockerCompose)
	//helpers.AddToFile(pathToDockerFile, DockerFile)

	// создаем конфиги для nginx
	inner_helpers.CreateNginx(pathToFrontServiceDockerDEV)
	// создаем конфиги для node
	inner_helpers.CreateNode(pathToFrontServiceDockerDEV)

	console.Info("UP NODE FOR INSTALL FRONTEND FRAMEWORK")

	// подымаем сервисы локально, чтоб развернуть framework для фронта
	console.Run("docker-compose", "-f", pathToLocalDockerComposeFile, "build")
	console.Run("docker-compose", "-f", pathToLocalDockerComposeFile, "up", "-d")

	// подымаем VUE
	if data.Driver == vue {
		// npm install -g @vue/cli
		// docker-compose run --rm front-nodejs vue create -d -f .
		// docker exec -it front_front_1 bash
		// os.Exit(2)

		// todo нужно реализовать создание проекта через пресеты (ща используеться установка по дефолту)
		// https://cli.vuejs.org/ru/guide/plugins-and-presets.html#%D0%BF%D1%80%D0%B5%D1%81%D0%B5%D1%82%D1%8B-%D0%BD%D0%B0%D1%81%D1%82%D1%80%D0%BE%D0%B5%D0%BA

		//vue create my-project
		console.Info("CREATE VUE PROJECT")
		console.Run("docker-compose", "-f", pathToLocalDockerComposeFile, "run", "--rm", "-u" , "node:node", "front", "vue", "create", "-d", "-f", tempName)

		// переносим vue в папку проекта
		helpers.MoveAllFromDir(pathToFrontService + "/" + tempName, pathToFrontService)
		helpers.DeleteDirectory(pathToFrontService + "/" + tempName)

		// билдим vue
		console.Info("BUILD VUE")
		console.Run("docker-compose", "-f", pathToLocalDockerComposeFile, "run", "--rm", "-u" , "node:node", "front", "yarn", "build")

		// останавливаем контейнер
		console.Info("DOWN CONTAINER")
		console.Run("docker-compose", "-f", pathToLocalDockerComposeFile, "down", "--remove-orphans")
	}

	return replaceDockerCompose(data.Name, projectName, ip)
}

func replaceDockerCompose(name string, pname string, ip string) string {

	dockerCompose = strings.Replace(dockerCompose, "{name}", name,10)
	dockerCompose = strings.Replace(dockerCompose, "{pname}", pname,4)
	dockerCompose = strings.Replace(dockerCompose, "{ip}", ip,1)

	return dockerCompose
}

