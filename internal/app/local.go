package app

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

type Local struct {
	Local string
}

const (
	ru = "ru"
	en = "en"
)
var locales = []string{ru, en}

func ChoiceLocal() *Local {

	result ,err := console.Select(console.NewDataForSelect("Choice language", locales))

	if err != nil {
		fmt.Printf("Prompt failed %v\n", result)
	}

	return &Local{
		Local: result,
	}
}
