package server

import (
	"github.com/romaxa83/skeleton/internal/helpers"
	"strings"
)

const (
	pathToDockerNginx = "/docker/dev/nginx"
	pathToConfNginx = pathToDockerNginx + "/conf.d"
)

var dockerFileNginx string = `FROM nginx:1.17-alpine

COPY ./dev/nginx/conf.d /etc/nginx/conf.d

WORKDIR /app
`
var conf string = `server {
    listen 80;
    charset utf-8;
    server_tokens off;
    index index.php;
    root /app/public;
    
    error_log  /var/log/nginx/error.log;
    access_log /var/log/nginx/access.log;

    location / {
        try_files $uri /index.php?$args;
    }

    location ~ \.php$ {
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_pass php-fpm:9000;
        fastcgi_index index.php;
        fastcgi_read_timeout 300;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param PATH_INFO $fastcgi_path_info;
    }
}
`

var serviceDockerComposeNginx string = `
  nginx:
    build:
      context: docker
      dockerfile: dev/nginx/Dockerfile
    container_name: {name}__nginx
    hostname: {name}__nginx
    volumes:
      - ./:/app
      - ./docker/storage/logs:/var/log/nginx
    ports:
      - {ip}:80:80
    restart: always
`

var serviceDockerComposeApache string = `
  apache:
    image: webdevops/apache:alpine
    container_name: {name}__apache
    hostname: {name}__apache
    environment:
      WEB_DOCUMENT_ROOT: /app/public
      WEB_PHP_SOCKET: php-fpm:9000
      LOG_STDOUT: /app/docker/storage/logs/web.access.log
      LOG_STDERR: /app/docker/storage/logs/web.errors.log
    volumes:
      - ./:/app:rw,cached
    ports:
      - {ip}:80:80
    working_dir: /app
    restart: always
    depends_on:
      - php-fpm
`

func Create(
	path string,
	driver string,
	ip string,
	projectName string,
) string {

	if driver == nginx {
		return createNginx(path, ip, projectName)
	}

	return createApache(ip, projectName)
}

func createNginx(
	path string,
	ip string,
	projectName string,
) string {
	pathToConfFileNginx := path + pathToConfNginx + "/default.conf"
	pathToDockerfileNginx := path + pathToDockerNginx + "/Dockerfile"

	helpers.CreateDir(path + pathToConfNginx)
	helpers.CreateAndWriteFile(pathToConfFileNginx, conf)
	helpers.CreateAndWriteFile(pathToDockerfileNginx, dockerFileNginx)

	t := strings.Replace(serviceDockerComposeNginx, "{ip}", ip,1)
	t = strings.Replace(t, "{name}", projectName,2)

	return t
}

func createApache(
	ip string,
	projectName string,
) string {

	t := strings.Replace(serviceDockerComposeApache, "{ip}", ip,1)
	t = strings.Replace(t, "{name}", projectName,2)

	return t
}


