package db

import (
	"github.com/romaxa83/skeleton/internal/helpers"
	"strings"
)

const (
	pathToStorage = "/docker/storage/db"
)

var serviceDockerComposeMysql string = `
  db:
    image: mysql:{version}
    container_name: {name}__db
    hostname: {name}__db
    restart: always
    environment:
      MYSQL_USER: {username}
      MYSQL_PASSWORD: {password}
      MYSQL_ROOT_USER: {username}
      MYSQL_ROOT_PASSWORD: {password}
      MYSQL_DATABASE: {db}
    ports:
      - {ip}:3306:3306
    volumes:
      - ./docker/storage/db/mysql:/var/lib/mysql
`

var serviceDockerComposePostgres string = `
  db:
    image: postgres:{version}
    container_name: {name}_db
    hostname: {name}_db
    restart: always
    environment:
      POSTGRES_USER: {username}
      POSTGRES_PASSWORD: {password}
      POSTGRES_DB: {db}
    ports:
      - {ip}:5432:5432
    volumes:
      - ./docker/storage/db/postgresql:/var/lib/postgresql/data
`

func Create(
	path string,
	driver string,
	ip string,
	projectName string,
	version string,
	username string,
	password string,
	dbName string,
) string {

	fullPathToStorage := path + pathToStorage
	helpers.CreateDir(fullPathToStorage)

	if driver == mysqlDriver {
		return createMysql(ip, projectName, version, username, password, dbName)
	}

	return createPostgres(ip, projectName, version, username, password, dbName)
}

func createMysql(
	ip string,
	projectName string,
	version string,
	username string,
	password string,
	dbName string,
) string {

	t := strings.Replace(serviceDockerComposeMysql, "{ip}", ip,1)
	t = strings.Replace(t, "{version}", version,1)
	t = strings.Replace(t, "{name}", projectName,2)
	t = strings.Replace(t, "{username}", username,2)
	t = strings.Replace(t, "{password}", password,2)
	t = strings.Replace(t, "{db}", dbName,2)

	return t
}

func createPostgres(
	ip string,
	projectName string,
	version string,
	username string,
	password string,
	dbName string,
) string {

	t := strings.Replace(serviceDockerComposePostgres, "{ip}", ip,1)
	t = strings.Replace(t, "{version}", version,2)
	t = strings.Replace(t, "{name}", projectName,2)
	t = strings.Replace(t, "{username}", username,2)
	t = strings.Replace(t, "{password}", password,2)
	t = strings.Replace(t, "{db}", dbName,2)

	return t
}
