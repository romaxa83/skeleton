package redis

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
	"github.com/romaxa83/skeleton/internal/helpers"
)

type Data struct {
	Install bool
	Password string
}

const (
	yes = "Yes"
	no = "No"
	title = "Install redis (for cache)"
	titlePassword = "Enter password for redis"
)
var ask = []string{yes, no}

func GetData() *Data {

	a, err := console.Select(console.NewDataForSelect(title, ask))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	an := helpers.ConvertBoolFromString(a)

	pass := ""
	if an {
		// запрашиваем пароль для redis
		p, err := console.Input(console.NewDataForInput(titlePassword, "secret"))
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}

		pass = p
	}

	return &Data{
		Install: an,
		Password: pass,
	}
}