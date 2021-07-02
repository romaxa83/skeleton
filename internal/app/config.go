package app

import (
	"github.com/romaxa83/skeleton/internal/entities/front_framework"
	"github.com/romaxa83/skeleton/internal/entities/ip"
	"github.com/romaxa83/skeleton/internal/entities/project_name"
	"github.com/romaxa83/skeleton/internal/entities/project_path"
)

//type Config struct {
//	ConfigMonolit ConfigMonolit
//	ConfigMicroservice ConfigMicroservice
//	ConfigOne ConfigOne
//}
//
//type ConfigMonolit struct {
//	ProjectName    *project_name.Data
//	ProjectPath    *project_path.Data
//	IP 			   *ip.Data
//	Server 		   *server.Data
//}
//
//type ConfigMicroservice struct {
//	ProjectName    *project_name.Data
//	ProjectPath    *project_path.Data
//	IP 			   *ip.Data
//	FrontFramework *front_framework.Data
//}
//
//type ConfigOne struct {
//	ProjectName    *project_name.Data
//	ProjectPath    *project_path.Data
//	IP 			   *ip.Data
//}
//
//func InitMonolitConfig() *ConfigMonolit  {
//	return &ConfigMonolit{
//		ProjectName:    project_name.GetData(),
//		ProjectPath:    project_path.GetData(),
//		IP:             ip.GetData(),
//		Server:         server.GetData(),
//	}
//}
//
//func InitMicroserviceConfig() *ConfigMicroservice {
//	return &ConfigMicroservice{
//		ProjectName:    project_name.GetData(),
//		ProjectPath:    project_path.GetData(),
//		IP:             ip.GetData(),
//		FrontFramework: front_framework.GetData(),
//	}
//}
//
//func InitOneConfig() *ConfigOne {
//	return &ConfigOne{
//		ProjectName:    project_name.GetData(),
//		ProjectPath:    project_path.GetData(),
//		IP:             ip.GetData(),
//	}
//}
//
//func InitConfig() *Config {
//	return &Config{
//		ConfigMonolit: InitMonolitConfig(),
//		ConfigMicroservice: InitMicroserviceConfig(),
//		ConfigOne: InitMicroserviceConfig(),
//	}
//}


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

func InitConfig() *Config {
	return &Config{
		ProjectName:    project_name.GetData(),
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
