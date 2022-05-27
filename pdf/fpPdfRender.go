package pdf

import (
	"fmt"
	"strconv"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

/*
	Format: ^FPd,g
	d =direction
		Values:
			H = horizontal printing (left to right)
			V = vertical printing (top to bottom)
			R = reverse printing (right to left)
	g = additional inter- character gap (in dots)
		Values:
			0 to 9999
			Default: 0 if no value is entered
*/
func fpRender(c zpl.Command, p *gopdf.GoPdf) *gopdf.GoPdf {
	params := c.GetParameters()

	if len(params) < 2 {
		return p
	}

	charterGap, err := strconv.ParseInt(params[1].Value, 4, 64);

	if nil != err {
		return p
	}

	fmt.Print(charterGap)

	switch direction := params[0].Value; direction {
	case "V":
		// Set vertical output options
		// https://kbpdfstudio.qoppa.com/adding-vertical-text-to-a-pdf-page/
	case "R":
		// Set reverse output options
		// https://answers.acrobatusers.com/Tiny-side-reversed-text-showing-up-q5414.aspx
		// https://en.wikipedia.org/wiki/Kerning
		// 
	default:
		// For H do nothing, this is default derection
		return p
	}

	return p
}
