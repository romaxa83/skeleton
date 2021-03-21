package app

import (
	"github.com/romaxa83/skeleton/internal/entities/db"
	"github.com/romaxa83/skeleton/internal/entities/framework"
	"github.com/romaxa83/skeleton/internal/entities/ip"
	"github.com/romaxa83/skeleton/internal/entities/mailer"
	"github.com/romaxa83/skeleton/internal/entities/nodejs"
	"github.com/romaxa83/skeleton/internal/entities/php"
	"github.com/romaxa83/skeleton/internal/entities/project_name"
	"github.com/romaxa83/skeleton/internal/entities/project_path"
	"github.com/romaxa83/skeleton/internal/entities/redis"
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
	Nodejs      *nodejs.Data
	Redis 		*redis.Data
	Mailer 		*mailer.Data
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
		Nodejs:      nodejs.GetData(),
		Redis:       redis.GetData(),
		Mailer:      mailer.GetData(),
	}
}
