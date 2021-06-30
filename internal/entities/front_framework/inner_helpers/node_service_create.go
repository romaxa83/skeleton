package inner_helpers

import "github.com/romaxa83/skeleton/internal/helpers"

var dockerFile string = `FROM node:latest

RUN yarn global add @vue/cli

USER node

WORKDIR /app
`

func CreateNode(
	pathToDockerFolder string,
) {
	pathToNode := pathToDockerFolder + "/node"
	pathToNodeDockerFile := pathToNode + "/Dockerfile"

	helpers.CreateDir(pathToNode)
	helpers.AddToFile(pathToNodeDockerFile, dockerFile)
}
