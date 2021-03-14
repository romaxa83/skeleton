package app

import (
	"github.com/romaxa83/skeleton/internal/entities/project_name"
	"github.com/romaxa83/skeleton/internal/entities/project_path"
)

type Config struct {
	projectName *project_name.Data
	projectPath *project_path.Data
	//server *server.Data
	//php *php.Data
	//db *db.Data
	//nodejs *nodejs.Data
	//redis *redis.Data
}

func InitConfig() *Config {
	return &Config{
		projectName: project_name.GetData(),
		projectPath: project_path.GetData(),
		//server:      server.GetData(),
		//php:         php.GetData(),
		//db:          db.GetData(),
		//nodejs:      nodejs.GetData(),
		//redis:       redis.GetData(),
	}
}
