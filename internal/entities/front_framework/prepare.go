package front_framework

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

const (
	vue = "Vue"
	react = "React (in developing)"
	useFramework = "Use Frontend Framework"
	titleDriver = "Choice Frontend Framework"
	name = "Name project"
)

var slimVersion = make([]string, 0)
var frameworkDriver = []string{vue, react}

type Data struct {
	UseFramework bool
	Driver string
	Name string
}

func GetData() *Data {

	// используем или нет фреймворк
	an := console.Ask(useFramework)
	if an {
		// название сервиса
		n, err := console.Input(console.NewDataForInput(name, "front"))
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}

		// запрашиваем тип фреймфорка
		d, err := console.Select(console.NewDataForSelect(titleDriver, frameworkDriver))
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}


		// todo добавить react
		// т.к. react  вразработке отключаем использование
		if d == react { an = false }

		return &Data{
			UseFramework: an,
			Driver: d,
			Name: n,
		}
	}

	return &Data{
		UseFramework: an,
	}
}