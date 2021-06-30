package back_framework

import (
	"github.com/romaxa83/skeleton/internal/console"
)

func Install(data *Data, path string, pathToDockerCompose string) {

	if data.Driver == laravel {
		createLaravel(data, path, pathToDockerCompose)
	}
	// todo реализация для slim

}

func createLaravel(data *Data, path string, pathToDockerCompose string)  {

	//laraPath := path + "/laravel/"

	console.Info("INSTALL LARAVEL")

	// устанавливаем laravel
	console.Run("docker-compose", "-f", pathToDockerCompose, "run", "--rm", "php-fpm", "composer", "create-project", "--prefer-dist", "laravel/laravel:^" + data.Version +".0")

	// переносим laravel в папку проекта
	//helpers.MoveAllFromDir(laraPath, path)
	//helpers.DeleteDirectory(laraPath)

	// останавливаем контейнеры чтоб перезаписать конфиги и поднять с новыми конфигами
	console.Info("STOPPED DOCKERS")
	console.Run("docker-compose", "-f", pathToDockerCompose , "down", "-v", "--remove-orphans")
}