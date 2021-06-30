package app

import (
	"github.com/romaxa83/skeleton/internal/entities/front_framework"
	"github.com/romaxa83/skeleton/internal/entities/ip"
	"github.com/romaxa83/skeleton/internal/entities/project_name"
	"github.com/romaxa83/skeleton/internal/entities/project_path"
)

type Config struct {
	ProjectName    *project_name.Data
	ProjectPath    *project_path.Data
	IP 			   *ip.Data
	//Arch 		   *arch.Data
	//Server 		   *server.Data
	//Php 		   *php.Data
	//DB 			   *db.Data
	//BackFramework  *back_framework.Data
	FrontFramework *front_framework.Data
	//Nodejs         *nodejs.Data
	//Redis 		   *redis.Data
	//Mailer 		   *mailer.Data
}

func InitConfig(local *Local) *Config {
	return &Config{
		ProjectName:    project_name.GetData(local.Local),
		ProjectPath:    project_path.GetData(),
		IP:             ip.GetData(),
		//Arch:           arch.GetData(),
		//Server:         server.GetData(),
		//Php:            php.GetData(),
		//DB:             db.GetData(),
		//BackFramework:  back_framework.GetData(),
		FrontFramework: front_framework.GetData(),
		//Nodejs:         nodejs.GetData(),
		//Redis:          redis.GetData(),
		//Mailer:         mailer.GetData(),
	}
}
