package project_name

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

const (
	title = "Enter project name"
)

type Data struct {
	Name string
}

func GetData() *Data {
	// запрашиваем название проекта
	n, err := console.Input(console.NewDataForInput(title, "test_project"))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return &Data{
		Name: n,
	}
}