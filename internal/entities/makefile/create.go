package makefile

import (
	"github.com/romaxa83/skeleton/internal/helpers"
	"strings"
)

var makeFile string = `
.SILENT:

#=============VARIABLES================
php_container = {name}__php-fpm
db_container = {name}__db
#======================================

#=====MAIN_COMMAND=====================

up: up_docker info

rebuild: down build up_docker info

up_docker:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

# флаг -v удаляет все volume (очищает все данные)
down-clear:
	docker-compose down -v --remove-orphans

build:
	docker-compose build

ps:
	docker-compose ps

#=======INTO_CONTAINER=================

php_bash:
	docker exec -it $(php_container) bash

db_bash:
	docker exec -it $(db_container) sh

#=======INTO_CONTAINER=================

info:
	echo "http://{ip}";
`

func Create(path string, projectName string, ip string) {

	makeFile = strings.Replace(makeFile, "{name}", projectName,2)
	makeFile = strings.Replace(makeFile, "{ip}", ip, 1)
	helpers.CreateAndWriteFile(path + "/Makefile", makeFile)
}
