package zpl_test

import (
	"testing"
	zpl "zplgun/zpl"
)

func TestFdParser(t *testing.T) {

	t.Run("test fdParser", func(t *testing.T) {
		zplData := zplData()
		commands := zpl.Parser([]byte(zplData))
		c := commands[14]
		params := c.GetParameters()

		if len(params) != 1 {
			t.Fatal("Got wrong number of params from buffer")
		}

		fieldText := params[0].Value
		expectedText := "Intershipping, Inc."

		if fieldText != expectedText {
			t.Fatal("Field text from params has to be Intershipping, Inc., got: " + fieldText)
		}
	})
}
