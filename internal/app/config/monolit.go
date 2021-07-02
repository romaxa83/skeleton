package config

import (
	"github.com/romaxa83/skeleton/internal/entities/ip"
	"github.com/romaxa83/skeleton/internal/entities/project_name"
	"github.com/romaxa83/skeleton/internal/entities/project_path"
	"github.com/romaxa83/skeleton/internal/entities/server"
)

type MonolitConfig struct {
	ProjectName    *project_name.Data
	ProjectPath    *project_path.Data
	IP 			   *ip.Data
	Server 		   *server.Data
}

func InitMonolitConfig() *MonolitConfig {
	return &MonolitConfig{
		ProjectName:    project_name.GetData(),
		ProjectPath:    project_path.GetData(),
		IP:             ip.GetData(),
		Server: 		server.GetData(),
	}
}
