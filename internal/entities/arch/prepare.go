package arch

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

type Data struct {
	Type string
}
const (
	title = "Choice app architecture"
	monolith = "monolith"
	microservice = "microservice"
)
var archType = []string{monolith, microservice}

func GetData() *Data {

	result ,err := console.Select(console.NewDataForSelect(title, archType))

	if err != nil {
		fmt.Printf("Prompt failed %v\n", result)
	}

	return &Data{
		Type: result,
	}
}