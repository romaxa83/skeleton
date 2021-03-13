package server

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

type Data struct {
	Name string
}

func GetData() *Data {

	n, err := choiceVersion()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return &Data{
		Name: n,
	}
}

func choiceVersion() (string, error) {
	prompt := promptui.Select{
		Label: "Choice Server",
		Size: 9,
		HideHelp: true,
		Items: []string{"Nginx", "Apache"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	fmt.Printf("You choose %q\n", result)

	return result, nil
}
