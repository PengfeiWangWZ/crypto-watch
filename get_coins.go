package main

import (
	"github.com/manifoldco/promptui"
)

func selectCoin(coins []Option) (int, error) {
	templates := &promptui.SelectTemplates{
		Label:    "Select a coin",
		Help:     "Use the arrow keys to navigate: ↓ ↑",
		Active:   "\U0001F4b0   {{ .Coin | cyan }}",
		Selected: "\U0001F4b0   {{ .Option | cyan }}",
		Inactive: "     {{ .Coin }}",
		Details:  "------End------",
	}

	prompt := promptui.Select{
		Items: coins,
		Templates: templates,
		Size: 10,
	}

	i, _, err := prompt.Run()
	return i, err
}