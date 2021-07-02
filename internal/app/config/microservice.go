package config

import (
	"github.com/romaxa83/skeleton/internal/entities/back_framework"
	"github.com/romaxa83/skeleton/internal/entities/db"
	"github.com/romaxa83/skeleton/internal/entities/front_framework"
	"github.com/romaxa83/skeleton/internal/entities/ip"
	"github.com/romaxa83/skeleton/internal/entities/php"
	"github.com/romaxa83/skeleton/internal/entities/project_name"
	"github.com/romaxa83/skeleton/internal/entities/project_path"
)

type MicroserviceConfig struct {
	ProjectName    *project_name.Data
	ProjectPath    *project_path.Data
	IP 			   *ip.Data
	FrontFramework *front_framework.Data
	Php 		   *php.Data
	DB 			   *db.Data
	BackFramework  *back_framework.Data
}

func InitMicroserviceConfig() *MicroserviceConfig {
	return &MicroserviceConfig{
		ProjectName:    project_name.GetData(),
		ProjectPath:    project_path.GetData(),
		IP:             ip.GetData(),
		FrontFramework: front_framework.GetData(),
		Php:            php.GetData(),
		DB:             db.GetData(),
		BackFramework:  back_framework.GetData(),
	}
}