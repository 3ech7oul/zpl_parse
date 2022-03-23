package zpl

import (
	"errors"
	"regexp"
)

type Command struct {
	ZplCommToken string
	Buffer       []byte
	parameters   []CommandParameter
}

type CommandParameter struct {
	Value string
}

func CreateCommand(commandToken string) (Command, error) {
	var result Command

	if isCommandToken(commandToken) {
		result = Command{
			ZplCommToken: commandToken,
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

func (c *Command) GetParameters() []CommandParameter {

	if 0 < len(c.parameters) {
		return c.parameters
	}

	handler := c.findHandler()
	if nil != handler {
		c.parameters = handler(c.ZplCommToken, c.Buffer)
	}

	return c.parameters
}

func (c *Command) findHandler() ParametersHandlerFunc {
	result, ok := ZplPdfHandlers[c.ZplCommToken]
	if ok {
		return result
	}

	return nil
}
