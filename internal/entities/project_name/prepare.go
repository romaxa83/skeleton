package project_name

import (
	"fmt"
	//"github.com/romaxa83/skeleton/internal/app"
	"github.com/romaxa83/skeleton/internal/console"
)

const (
	titleRu = "Введите название проекта"
	titleEn = "Enter project name"
)

type Data struct {
	Name string
}

func GetData(local string) *Data {
	//fmt.Println(local)
	//fmt.Println(app.TitleProjectMame)

	// запрашиваем название проекта
	n, err := console.Input(console.NewDataForInput(titleRu, "test_project"))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return &Data{
		Name: n,
	}
}