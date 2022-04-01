package zpl_test

import (
	"testing"
	zpl "zplgun/zpl"
)

func TestFdParser(t *testing.T) {

	t.Run("test fdParser", func(t *testing.T) {
		zplData := zplData()
		commands := zpl.Parser([]byte(zplData))
		c := commands[256]
		params := c.GetParameters()

		if len(params) != 1 {
			t.Fatal("Got wrong number of params from buffer")
		}

		fieldText := params[0].Value
		expectedText := "Shelbyville TN 38102"

		if fieldText != expectedText {
			t.Fatal("Field text from params has to be Shelbyville TN 38102, got: " + fieldText)
		}
	})
}
