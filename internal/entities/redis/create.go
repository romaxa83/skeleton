package redis

import (
	"strings"
)

var dockerCompose string = `
  redis:
    image: redis:5.0-alpine
    container_name: {name}__redis
    hostname: {name}__redis
    ports:
      - {ip}:6379:6379
    command:
      - 'redis-server'
      - '--databases 2'
      - '--save 900 1'
      - '--save 300 10'
      - '--save 60 10000'
      - '--requirepass {password}'
`

func Create(data *Data, path string, projectName string, ip string) string {

	dockerCompose = strings.Replace(dockerCompose, "{name}", projectName,-1)
	dockerCompose = strings.Replace(dockerCompose, "{ip}", ip,1)
	dockerCompose = strings.Replace(dockerCompose, "{password}", data.Password,1)

	return dockerCompose
}