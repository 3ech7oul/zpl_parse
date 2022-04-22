package zpl

import "fmt"

func Parser(input []byte) []Command {
	var buffCommands []Command
	var currentCommand Command

	for i, s := range input {
		var commandToken string

		if string(s) == "^" {
			commandToken = fmt.Sprintf("^%s%s", string(input[i+1]), string(input[i+2]))
		}

		c, err := CreateCommand(commandToken)
		if nil == err {
			if 0 < len(currentCommand.ZplCommToken) {
				buffCommands = append(buffCommands, currentCommand)
			}
			currentCommand = c
		} else {
			currentCommand.AddToBuffer(s)
		}
	}

	return buffCommands
}
