package pkg

import "github.com/charmbracelet/huh"

func AskForInputBlocking(prompt string) string {
	result := ""
	huh.NewInput().
		Title(prompt).
		Prompt("> ").
		Value(&result).
		Run()
	return result
}
