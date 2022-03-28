package zpl

import "fmt"

func Parser(input []byte) map[int]Command {
	var buffCommands = make(map[int]Command)
	var currentCommand Command

	for i, s := range input {
		var commandToken string

		if "^" == string(s) {
			commandToken = fmt.Sprintf("^%s%s", string(input[i+1]), string(input[i+2]))
		}

		c, err := CreateCommand(commandToken)
		if nil == err {
			if 0 < len(currentCommand.ZplCommToken) {
				buffCommands[i] = currentCommand
			}
			currentCommand = c
		} else {
			currentCommand.AddToBuffer(s)
		}
	}

	return buffCommands
}
