package mailer

import (
	"github.com/romaxa83/skeleton/internal/console"
)

type Data struct {
	Install bool
}

const (
	title = "Install mailer"
)

func GetData() *Data {

	an := console.Ask(title)

	return &Data{
		Install: an,
	}
}
