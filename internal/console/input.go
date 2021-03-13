package console

import (
	"errors"
	"github.com/manifoldco/promptui"
	//"os/user"
)

type DataForInput struct {
	title string
	defaultValue string
}

func NewDataForInput(title string, defaultValue string) *DataForInput {
	return &DataForInput{
		title: title,
		defaultValue: defaultValue,
	}
}

func Input(d *DataForInput) (string, error) {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("Value must have more than 3 characters")
		}
		return nil
	}

	//var username string
	//u, err := user.Current()
	//if err == nil {
	//	username = u.Username
	//}

	prompt := promptui.Prompt{
		Label:    d.title,
		Validate: validate,
		Default:  d.defaultValue,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}