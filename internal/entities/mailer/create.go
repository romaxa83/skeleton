package mailer

import (
"strings"
)

var dockerCompose string = `
  mailer:
    image: mailhog/mailhog
    container_name: {name}__mailer
    hostname: {name}__mailer
    ports:
      - {ip}:8025:8025
      - {ip}:1025:1025
`

func Create(projectName string, ip string) string {

	dockerCompose = strings.Replace(dockerCompose, "{name}", projectName,-1)
	dockerCompose = strings.Replace(dockerCompose, "{ip}", ip,2)

	return dockerCompose
}
