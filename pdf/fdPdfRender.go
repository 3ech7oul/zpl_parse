package pdf

import (
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func fdRenderer(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf {
	params := c.GetParameters()

	if len(params) < 1 {
		return p
	}

	p.Cell(nil, params[0].Value)

	return p
}
