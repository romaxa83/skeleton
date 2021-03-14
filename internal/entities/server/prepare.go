package server

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

type Data struct {
	Name string
}
const (
	title = "Choice Server"
	nginx = "Nginx"
	apache = "Apache"
)
var serverName = []string{nginx, apache}

func GetData() *Data {

	result ,err := console.Select(console.NewDataForSelect(title, serverName))

	if err != nil {
		fmt.Printf("Prompt failed %v\n", result)
	}

	return &Data{
		Name: result,
	}
}
