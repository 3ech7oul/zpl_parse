package zpl

import "fmt"

func Parese(input []byte) {
	var buffCommands = make(map[int]Command)
	var currentCommand Command

	for i, item := range input {
		var commandToken string

		if "^" == string(item) {
			commandToken = fmt.Sprintf("^%s%s", string(input[i+1]), string(input[i+2]))
		}

		comm, err := CreateCommand(commandToken)

		if nil == err {
			if 0 < len(currentCommand.ZplComm) {
				buffCommands[i] = currentCommand
			}
			currentCommand = comm
		} else {
			currentCommand.AddToBuffer(string(item))
		}
	}

	fmt.Print(buffCommands)

	//fmt.Print(currentCommand)

}
