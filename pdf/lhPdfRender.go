package pdf

import (
	"regexp"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func lhRender(c zpl.Command, p *Pdf) *gopdf.GoPdf {
	params := c.GetParameters()

	if len(params) < 2 {
		return &p.pdf
	}

	nonStrRE := regexp.MustCompile(`\D*`)
	p.globalSettings["labelHomeX"] = nonStrRE.ReplaceAllString(params[0], "")
	p.globalSettings["labelHomeY"] = nonStrRE.ReplaceAllString(params[1], "")
	return &p.pdf
}
