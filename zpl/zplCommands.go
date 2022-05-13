package zpl

import (
	"errors"
	"regexp"
	"strings"
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

	handler := c.findParser()
	if nil != handler {
		v := removeZplCommandToken(string(c.Buffer), c.ZplCommToken)
		params := strings.Split(v, ",")

		c.parameters = handler(params)
	}

	return c.parameters
}

func (c *Command) findParser() ParametersParserFunc {
	result, ok := ZplPdfParsers[c.ZplCommToken]
	if ok {
		return result
	}

	return nil
}
