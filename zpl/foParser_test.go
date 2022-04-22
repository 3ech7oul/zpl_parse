package zpl_test

import (
	"testing"
	zpl "zplgun/zpl"
)

func TestFoParser(t *testing.T) {

	t.Run("test foParser", func(t *testing.T) {
		zplData := zplData()
		commands := zpl.Parser([]byte(zplData))
		c := commands[3]
		params := c.GetParameters()

		if len(params) != 2 {
			t.Fatal("Got wrong number of params from buffer")
		}

		x := params[0].Value
		y := params[1].Value
		expectedX := "50"
		expectedY := "50"

		if x != expectedX {
			t.Fatal("X coord from params has to be 100, got: " + x)
		}

		if y != expectedY {
			t.Fatal("Y coord from params has to be 960, got: " + y)
		}
	})
}
