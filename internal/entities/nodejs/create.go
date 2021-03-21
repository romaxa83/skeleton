package nodejs

import (
	"github.com/romaxa83/skeleton/internal/helpers"
	"strings"
)

const (
	pathToDockerNode = "/docker/dev/node"
)

var dockerFile string = `FROM node:{version}-alpine

WORKDIR /app
`

var dockerCompose string = `
  node:
    build:
      context: docker
      dockerfile: dev/node/Dockerfile
    container_name: {name}__node
    hostname: {name}__node
    volumes:
      - ./:/app
    tty: true
`

func Create(data *Data, path string, projectName string) string {

	pathToConf := path + pathToDockerNode
	helpers.CreateDir(pathToConf)

	dockerFile = strings.Replace(dockerFile, "{version}", data.Version,1)
	helpers.CreateAndWriteFile(pathToConf + "/Dockerfile", dockerFile)

	dockerCompose = strings.Replace(dockerCompose, "{name}", projectName,-1)

	return dockerCompose
}
