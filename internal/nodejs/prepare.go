package nodejs

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
	"github.com/romaxa83/skeleton/internal/helpers"
)

type Data struct {
	Install bool
	Version string
}

const (
	yes = "Yes"
	no = "No"
	title = "Install nodejs"
	titleVersion = "Choice Nodejs version"
)
var ask = []string{yes, no}
var versions = []string{"13", "12", "8"}

func GetData() *Data {

	a, err := console.Select(console.NewDataForSelect(title, ask))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	an := helpers.ConvertBoolFromString(a)
	v := ""

	if an {
		ver, err := console.Select(console.NewDataForSelect(titleVersion, versions))
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}

		v = ver
	}

	return &Data{
		Install: an,
		Version: v,
	}
}
