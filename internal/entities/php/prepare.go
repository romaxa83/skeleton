package php

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

type Data struct {
	Version string
}

const (
	title = "Choice Php version"
)
var version = []string{"8.0", "7.4", "7.3"}

func GetData() *Data {

	v, err := console.Select(console.NewDataForSelect(title, version))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return &Data{
		Version: v,
	}
}