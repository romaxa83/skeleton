package app

import (
	"github.com/romaxa83/skeleton/internal/helpers"
)

func Create(config *Config) {

	path := config.projectPath.Path + "/" + config.projectName.Name

	helpers.CreateDir(path)

	//fmt.Println(path)
}
