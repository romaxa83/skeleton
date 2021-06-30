package back_framework

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

const (
	laravel = "Laravel"
	slim = "Slim (in developing)"
	useFramework = "Use Backend Framework"
	titleDriver = "Choice Backend Framework"
	titleVersion = "Choice Backend Framework version"
	titleRunMigration = "Run migration"
)

var laravelVersion = []string{"8", "7"}
var slimVersion = make([]string, 0)
var frameworkDriver = []string{laravel, slim}

type Data struct {
	UseFramework bool
	Driver string
	Version string
	RunMigration bool
	RunSeed bool
}

func GetData() *Data {

	// используем фреймфорк
	an := console.Ask(useFramework)
	if an {
		// запрашиваем тип фреймфорка
		d, err := console.Select(console.NewDataForSelect(titleDriver, frameworkDriver))
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}

		// запрашиваем версию фреймфорка
		versions := getVersionFramework(d)
		v, err := console.Select(console.NewDataForSelect(titleVersion, versions))
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}

		m := console.Ask(titleRunMigration)

		// todo добавить slim
		// т.к. slim  вразработке отключаем использование
		if d == slim { an = false }

		return &Data{
			UseFramework: an,
			Driver: d,
			Version: v,
			RunMigration: m,
		}
	}

	return &Data{
		UseFramework: an,
	}
}

func getVersionFramework(driver string) []string {
	if driver == laravel {
		return laravelVersion
	}

	return slimVersion
}
