package console

import "github.com/manifoldco/promptui"

type DataForSelect struct {
	title string
	items []string
}

func NewDataForSelect(title string, items []string) *DataForSelect {
	return &DataForSelect{
		title: title,
		items: items,
	}
}

func Select(d *DataForSelect) (string, error) {
	prompt := promptui.Select{
		Label: d.title,
		Size: 9,
		HideHelp: true,
		Items: d.items,
	}
	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

