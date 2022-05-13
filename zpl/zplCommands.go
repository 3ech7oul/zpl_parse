package zpl

import (
	"errors"
	"regexp"
)

type Command struct {
	ZplCommToken string
	Buffer       []byte
	parameters   []string
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

	return len(re.FindAllString(str, -1)) > 0
}

func (c *Command) AddToBuffer(b byte) {
	c.Buffer = append(c.Buffer, b)
}

func (c *Command) GetParameters() []string {

	if 0 < len(c.parameters) {
		return c.parameters
	}

	handler := c.findParser()
	if nil != handler {
		c.parameters = handler(c.ZplCommToken, c.Buffer)
	}

	return c.parameters
}

func (c *Command) findParser() ParametersParserFunc {
	result, ok := ZplPdfParsers[c.ZplCommToken]
	if ok {
		return result
	}

	return defaultParser
}
