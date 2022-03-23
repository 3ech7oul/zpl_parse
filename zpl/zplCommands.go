package zpl

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/signintech/gopdf"
)

var zplCommnads = map[string]ZplHandlerFunc{
	//"^FD": zplCfHandler,
	//"^FX": zplCfHandler,
	//"^XA": zplCfHandler,
	//"^FS": zplCfHandler,
	//"^GB": zplCfHandler,
	//"^FO": zplCfHandler,
	"^CF": zplCfHandler,
}

type ZplHandlerFunc func(c Command)

type Command struct {
	ZplComm    string
	Buffer     []byte
	Parameters []CommandParameter
}

type CommandParameter struct {
	Index int
	Value string
}

func CreateCommand(commandToken string) (Command, error) {
	var result Command

	if isCommandToken(commandToken) {
		result = Command{
			ZplComm: commandToken,
		}

		return result, nil
	}

	return result, errors.New("Not a command")
}

func isCommandToken(str string) bool {
	var re = regexp.MustCompile(`(?m)(\^)+(\D)+(\D)`)
	if len(re.FindAllString(str, -1)) > 0 {
		return true
	}

	return false
}

func (c *Command) AddToBuffer(b byte) {
	c.Buffer = append(c.Buffer, b)
}

func (c *Command) GetParameters() {
	fmt.Println(string(c.Buffer))
}

func zplCfHandler(c Command) {
	pdf := gopdf.GoPdf{}
	fmt.Println(pdf)
}
