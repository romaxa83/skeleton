package inner_helpers

import "github.com/romaxa83/skeleton/internal/helpers"

var nginxDockerFile string = `FROM nginx:1.17-alpine

COPY ./nginx/conf.d /etc/nginx/conf.d

WORKDIR /app
`

var nginxConf string = `server {
    listen 80;
    index index.html;
    root /app/dist;

    location / {
        try_files $uri /index.html;
    }
}
`
//создаем конфиги для nginx

func CreateNginx(
	pathToDockerFolder string,
) {

	pathToNginx := pathToDockerFolder + "/nginx"
	pathToNginxDockerFile := pathToNginx + "/Dockerfile"
	pathToNginxConf := pathToNginx + "/conf.d"
	pathToNginxConfFile := pathToNginxConf + "/default.conf"


	helpers.CreateDir(pathToNginxConf)
	helpers.AddToFile(pathToNginxDockerFile, nginxDockerFile)
	helpers.AddToFile(pathToNginxConfFile, nginxConf)
}
