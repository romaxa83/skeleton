package app

import (
	"bufio"
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
	"github.com/romaxa83/skeleton/internal/entities/db"
	"github.com/romaxa83/skeleton/internal/entities/framework"
	"github.com/romaxa83/skeleton/internal/entities/mailer"
	"github.com/romaxa83/skeleton/internal/entities/makefile"
	"github.com/romaxa83/skeleton/internal/entities/nodejs"
	"github.com/romaxa83/skeleton/internal/entities/php"
	"github.com/romaxa83/skeleton/internal/entities/redis"
	"github.com/romaxa83/skeleton/internal/entities/server"
	"github.com/romaxa83/skeleton/internal/helpers"
	"io"
	"time"

	//"os"
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

func replace(r io.Reader, w io.Writer) error {
	// use scanner to read line by line
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		fmt.Println("LINE")
		fmt.Println(line)
		fmt.Println("---------------------------------")
		if line == "nginx" {
			line = strings.Replace(line, "nginx", "test",2)
			//line = "APP_NAME=TEST"
		}
		if _, err := io.WriteString(w, line+"\n"); err != nil {
			return err
		}
	}
	return sc.Err()
}

func Create(config *Config) {

	path := config.ProjectPath.Path + "/" + config.ProjectName.Name

	pathToDockerCompose := path + "/docker-compose.yml"

	// создаем папку проекта
	helpers.CreateDir(path)
	//создаем docker-compose с выбраными сервисами, а также создаем под них конфиги
	createDockerCompose(config, path, pathToDockerCompose)
	// создаем Makefile
	makefile.Create(path, config.ProjectName.Name, config.IP.IP)
	// подымаем контейнеры
	console.Run("docker-compose", "-f", pathToDockerCompose, "build")
	console.Run("docker-compose", "-f", pathToDockerCompose, "up", "-d")

	// запускаем установку framework если он выбран
	if config.Framework.UseFramework {
		framework.Install(config.Framework, path, pathToDockerCompose)
		// правим конфиг .env
		envFile := path + "/.env"
		helpers.ReplaceIntoFile(envFile, "APP_NAME=Laravel","APP_NAME=" + strings.Title(config.ProjectName.Name), 1)
		helpers.ReplaceIntoFile(envFile, "APP_URL=http://localhost", "APP_URL=http://" + config.IP.IP, 1)
		helpers.ReplaceIntoFile(envFile, "DB_CONNECTION=mysql", "DB_CONNECTION=" + config.DB.Driver, 1)
		helpers.ReplaceIntoFile(envFile, "DB_HOST=127.0.0.1", "DB_HOST=" + config.IP.IP, 1)
		helpers.ReplaceIntoFile(envFile, "DB_DATABASE=laravel", "DB_DATABASE=" + config.DB.DbName, 1)
		helpers.ReplaceIntoFile(envFile, "DB_USERNAME=root", "DB_USERNAME=" + config.DB.UserName, 1)
		helpers.ReplaceIntoFile(envFile, "DB_PASSWORD=", "DB_PASSWORD=" + config.DB.Password, 1)

		if config.Redis.Install {
			helpers.ReplaceIntoFile(envFile, "REDIS_HOST=127.0.0.1", "REDIS_HOST=" + config.IP.IP, 1)
			helpers.ReplaceIntoFile(envFile, "REDIS_PASSWORD=null", "REDIS_PASSWORD=" + config.Redis.Password, 1)
		}

		if config.Mailer.Install {
			helpers.ReplaceIntoFile(envFile, "REDIS_HOST=127.0.0.1", "REDIS_HOST=" + config.IP.IP, 1)
			helpers.ReplaceIntoFile(envFile, "REDIS_PASSWORD=null", "REDIS_PASSWORD=secret" + config.Redis.Password, 1)
		}

		helpers.AddToFile(envFile, "\nDOCKER_BRIDGE=" + config.IP.IP)
		helpers.AddToFile(envFile, "\nDOCKER_NETWORK=" + helpers.NetworkIP(config.IP.IP))
		// правим docker-compose
		helpers.ReplaceIntoFile(pathToDockerCompose, config.ProjectName.Name + "__", "${APP_NAME}__", -1)
		helpers.ReplaceIntoFile(pathToDockerCompose, config.IP.IP , "${DOCKER_BRIDGE}", -1)
		helpers.ReplaceIntoFile(pathToDockerCompose, helpers.NetworkIP(config.IP.IP) , "${DOCKER_NETWORK}", -1)
		helpers.ReplaceIntoFile(pathToDockerCompose, "MYSQL_USER: " + config.DB.UserName , "MYSQL_USER: ${DB_USERNAME}", 1)
		helpers.ReplaceIntoFile(pathToDockerCompose, "MYSQL_ROOT_USER: " + config.DB.UserName , "MYSQL_ROOT_USER: ${DB_USERNAME}", 1)
		helpers.ReplaceIntoFile(pathToDockerCompose, "MYSQL_PASSWORD: " + config.DB.Password , "MYSQL_PASSWORD: ${DB_PASSWORD}", 1)
		helpers.ReplaceIntoFile(pathToDockerCompose, "MYSQL_ROOT_PASSWORD: " + config.DB.Password , "MYSQL_ROOT_PASSWORD: ${DB_PASSWORD}", 1)
		helpers.ReplaceIntoFile(pathToDockerCompose, "MYSQL_DATABASE: " + config.DB.DbName , "MYSQL_DATABASE: ${DB_DATABASE}", 1)
		helpers.ReplaceIntoFile(pathToDockerCompose, "POSTGRES_PASSWORD: " + config.DB.Password , "POSTGRES_PASSWORD: ${DB_PASSWORD}", 1)
		helpers.ReplaceIntoFile(pathToDockerCompose, "POSTGRES_USER: " + config.DB.UserName , "POSTGRES_USER: ${DB_USERNAME}", 1)
		helpers.ReplaceIntoFile(pathToDockerCompose, "POSTGRES_DB: " + config.DB.DbName , "POSTGRES_DB: ${DB_DATABASE}", 1)
		// правим Makefile
		makefilePath := path + "/Makefile"
		helpers.ReplaceIntoFile(makefilePath, ".SILENT:", ".SILENT:\ninclude .env", 1)
		helpers.ReplaceIntoFile(makefilePath, "http://" + config.IP.IP, "${APP_URL}", 1)
		helpers.ReplaceIntoFile(makefilePath, config.ProjectName.Name + "__php-fpm", strings.Title(config.ProjectName.Name) + "__php-fpm", 1)
		helpers.ReplaceIntoFile(makefilePath, config.ProjectName.Name + "__db", strings.Title(config.ProjectName.Name) + "__db", 1)
		// запускаем контейнер
		console.Run("docker-compose", "-f", pathToDockerCompose, "--env-file", envFile, "up", "-d")
		time.Sleep(8 * time.Second)
		if config.Framework.RunMigration {
			console.Run("docker-compose", "-f", pathToDockerCompose, "--env-file", envFile, "run", "--rm", "php-fpm", "php", "artisan", "migrate")
		}
		if config.Framework.RunSeed {
			console.Run("docker-compose", "-f", pathToDockerCompose, "--env-file", envFile, "run", "--rm", "php-fpm", "php", "artisan", "db:seed")
		}
	}

	console.Success("http://" + config.IP.IP)
}

func createDockerCompose(config *Config, path string, pathToDockerCompose string)  {

	// создаем файл docker-compose
	helpers.AddToFile(pathToDockerCompose, startDockerCompose)

	// создаем сервер и прописываем docker-compose
	serviceServer := server.Create(path, config.Server.Name, config.IP.IP, config.ProjectName.Name)
	helpers.AddToFile(pathToDockerCompose, serviceServer)

	// создаем php и прописываем docker-compose
	servicePhp := php.Create(path, config.Php.Version, config.ProjectName.Name)
	helpers.AddToFile(pathToDockerCompose, servicePhp)

	// создаем nodeJS
	if config.Nodejs.Install {
		serviceNode := nodejs.Create(config.Nodejs, path, config.ProjectName.Name)
		helpers.AddToFile(pathToDockerCompose, serviceNode)
	}

	// создаем redis
	if config.Redis.Install {
		serviceRedis := redis.Create(config.Redis, path, config.ProjectName.Name, config.IP.IP)
		helpers.AddToFile(pathToDockerCompose, serviceRedis)
	}

	// создаем mailer
	if config.Mailer.Install {
		serviceMailer := mailer.Create(config.ProjectName.Name, config.IP.IP)
		helpers.AddToFile(pathToDockerCompose, serviceMailer)
	}

	// todo отрефакторить
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
