package php

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

type Data struct {
	Version string
}

func GetData() *Data {

	v, err := choiceVersion()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return &Data{
		Version: v,
	}
}

func choiceVersion() (string, error) {
	prompt := promptui.Select{
		Label: "Choice PHP version",
		Size: 9,
		HideHelp: true,
		Items: []string{"7.3", "7.4", "8.0"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	fmt.Printf("You choose %q\n", result)

	return result, nil
}