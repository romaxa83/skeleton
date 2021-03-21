package helpers

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
			DeleteDirectory(path + "/")
			createDirectory(path)
		} else {
			console.Error("stopped build")
			os.Exit(1)
		}

	} else {
		createDirectory(path)
	}
}

func CreateAndWriteFile(path string, text string)  {
	file, err := os.Create(path)

	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString(text)

	console.Success("create and write file - " + path)
}

// меняет вхожденмя в файле
func ReplaceIntoFile(path string, search string, replace string, n int)  {
	fContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	newFile := strings.Replace(string(fContent), search, replace,n)
	CreateAndWriteFile(path, newFile)
}

func AddToFile(path string, text string) {

	file, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Problem open file ("+ path +") :", err)
	}
	defer file.Close()

	if _, err := file.WriteString(text); err != nil {
		fmt.Println("Problem with add text to file ("+ path +") :", err)
	}

	console.Success("recorded to file - " + path)
}

func createDirectory(path string)  {
	os.MkdirAll(path, os.ModePerm)
	console.Success("create directory - " + path)
}

func DeleteDirectory(path string) {

	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	console.Success("deleted directory - " + path)
}

// перемещает содержимое из одной директории в другую
func MoveAllFromDir(path string, moveToPath string)  {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		// Обход директории без вывода
		if wPath == path {
			return nil
		}
		// Если данный путь является директорией, то останавливаем рекурсивный обход
		// и возвращаем название папки
		if info.IsDir() {
			console.Info("move directory - " + wPath)
			console.Run("mv", wPath, moveToPath)
			return filepath.SkipDir
		}
		// Выводится название файла
		if wPath != path {
			console.Info("move file - " + wPath)
			console.Run("mv", wPath, moveToPath)
		}
		return nil
	})
}
