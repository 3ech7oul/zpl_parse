package pdf

import (
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func fdRender(c zpl.Command, p *Pdf) *gopdf.GoPdf {
	params := c.GetParameters()

	if len(params) < 1 {
		return &p.pdf
	}

	p.Cell(nil, params[0])

	return &p.pdf
}
