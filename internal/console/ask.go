package console

import (
	"fmt"
)

const (
	yes = "Yes"
	no = "No"
)

var ask = []string{yes, no}

func Ask(title string) bool {
	a, err := Select(NewDataForSelect(title, ask))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	if a == yes {
		return true
	}

	return false
}