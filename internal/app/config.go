package app

import (
	"github.com/romaxa83/skeleton/internal/entities/ip"
	"github.com/romaxa83/skeleton/internal/entities/php"
	"github.com/romaxa83/skeleton/internal/entities/project_name"
	"github.com/romaxa83/skeleton/internal/entities/project_path"
	"github.com/romaxa83/skeleton/internal/entities/server"
)

type Config struct {
	ProjectName *project_name.Data
	ProjectPath *project_path.Data
	IP *ip.Data
	Server *server.Data
	Php *php.Data
	//db *db.Data
	//nodejs *nodejs.Data
	//redis *redis.Data
}

func InitConfig() *Config {
	return &Config{
		ProjectName: project_name.GetData(),
		ProjectPath: project_path.GetData(),
		IP: ip.GetData(),
		Server:      server.GetData(),
		Php:         php.GetData(),
		//db:          db.GetData(),
		//nodejs:      nodejs.GetData(),
		//redis:       redis.GetData(),
	}
}
