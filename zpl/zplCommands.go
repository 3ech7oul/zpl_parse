package zpl

import (
	"errors"
	"fmt"

	"github.com/signintech/gopdf"
)

var zplCommnads = map[string]ZplHandlerFunc{
	//"^FD": "",
	//"^FX": "",
	//"^XA": "",
	//"^FS": "",
	//"^GB": "",
	//"^FO": "",
	"^CF": zplCfHandler,
}

type ZplHandlerFunc func(c Command)

type Command struct {
	ZplComm    string
	Buffer     []string
	Parameters []CommandParameter
}

type CommandParameter struct {
	Index int
	Value string
}

func CreateCommand(command string) (Command, error) {
	var result Command

	_, ok := zplCommnads[command]
	if ok {
		result = Command{
			ZplComm: command,
		}

		return result, nil
	}

	return result, errors.New("Not a command")
}

func (c *Command) AddToBuffer(b string) {
	c.Buffer = append(c.Buffer, b)
}

func zplCfHandler(c Command) {
	pdf := gopdf.GoPdf{}
	fmt.Println(pdf)
}
