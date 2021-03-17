package app

import (
	"github.com/romaxa83/skeleton/internal/entities/db"
	"github.com/romaxa83/skeleton/internal/entities/framework"
	"github.com/romaxa83/skeleton/internal/entities/ip"
	"github.com/romaxa83/skeleton/internal/entities/php"
	"github.com/romaxa83/skeleton/internal/entities/project_name"
	"github.com/romaxa83/skeleton/internal/entities/project_path"
	"github.com/romaxa83/skeleton/internal/entities/server"
)

type Config struct {
	ProjectName *project_name.Data
	ProjectPath *project_path.Data
	IP 			*ip.Data
	Server 		*server.Data
	Php 		*php.Data
	DB 			*db.Data
	Framework 	*framework.Data
	//nodejs *nodejs.Data
	//redis *redis.Data
}

func InitConfig() *Config {
	return &Config{
		ProjectName: project_name.GetData(),
		ProjectPath: project_path.GetData(),
		IP:          ip.GetData(),
		Server:      server.GetData(),
		Php:         php.GetData(),
		DB:          db.GetData(),
		Framework:   framework.GetData(),
		//nodejs:      nodejs.GetData(),
		//redis:       redis.GetData(),
	}
}
