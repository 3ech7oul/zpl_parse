package pdf

import (
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func lhRender(c zpl.Command, p *Pdf) *gopdf.GoPdf {
	params := c.GetParameters()

	if len(params) < 2 {
		return &p.pdf
	}

	p.globalSettings["labelHomeX"] = params[0].Value
	p.globalSettings["labelHomeY"] = params[1].Value

	return &p.pdf
}
