package app

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

type Strategy struct {
	Type string
}

const (
	mono = "monolit"
	micro = "microservice"
	one = "one service (developing)"
)
var strategies = []string{mono, micro, one}

func ChoiceStrategy() *Strategy {

	result ,err := console.Select(console.NewDataForSelect("Choice strategy for app", strategies))

	if err != nil {
		fmt.Printf("Prompt failed %v\n", result)
	}

	return &Strategy{
		Type: result,
	}
}

func IsMonolit(strategy *Strategy) bool {
	return strategy.Type == mono
}

func IsMicroservice(strategy *Strategy) bool {
	return strategy.Type == micro
}

func IsOne(strategy *Strategy) bool {
	return strategy.Type == one
}