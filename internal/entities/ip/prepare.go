package ip

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

const (
	title = "Enter ip for docker network"
)

type Data struct {
	IP string
}

func GetData() *Data {
	// запрашиваем ip для сети докера
	ip, err := console.Input(console.NewDataForInput(title, "192.168.180.1"))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return &Data{
		IP: ip,
	}
}
