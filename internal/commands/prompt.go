package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	label    string
	errorMsg string
}

func promptCommand(pc promptContent, obscure ...bool) string {
	isObscure := false
	if len(obscure) > 0 {
		isObscure = obscure[0]
	}

	validate := func(input string) error {
		if strings.Trim(input, " ") == "" {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	var prompt promptui.Prompt

	if isObscure {
		prompt = promptui.Prompt{
			Label:    pc.label,
			Validate: validate,
			Mask:     '*',
		}
	} else {
		prompt = promptui.Prompt{
			Label:    pc.label,
			Validate: validate,
		}
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Println("Prompt failed", err)
		os.Exit(1)
	}

	return result
}
