package project_path

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
	"os"
)

const (
	title = "Enter project path"
)

type Data struct {
	Path string
}

func GetData() *Data {
	// получаем папку где запущена команда
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// запрашиваем к проекта
	n, err := console.Input(console.NewDataForInput(title, pwd))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return &Data{
		Path: n,
	}
}
