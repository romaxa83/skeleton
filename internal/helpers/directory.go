package helpers

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
	"os"
)

func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }

	return false, err
}

func CreateDir(path string)  {

	ex, err := Exist(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if ex {
		console.Warning("directory exist - " + path)

		an := console.Ask("Delete directory")

		if an {
			deleteDirectory(path + "/")
			createDirectory(path)
		} else {
			console.Error("stopped build")
			os.Exit(1)
		}

	} else {
		createDirectory(path)
	}
}

func createDirectory(path string)  {
	os.MkdirAll(path, os.ModePerm)
	console.Success("create directory - " + path)
}

func deleteDirectory(path string) {

	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	console.Success("deleted directory - " + path)
}
