package helpers

var serviceDockerComposePostgres string = `*
!.gitignore
`

func CreateIgnoreAll(path string)  {

	CreateAndWriteFile(path + "/.gitignore", serviceDockerComposePostgres)
}
