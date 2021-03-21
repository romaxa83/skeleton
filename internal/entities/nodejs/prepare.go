package nodejs

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

type Data struct {
	Install bool
	Version string
}

const (
	title = "Install nodejs"
	titleVersion = "Choice Nodejs version"
)

var versions = []string{"13", "12", "8"}

func GetData() *Data {

	an := console.Ask(title)
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
