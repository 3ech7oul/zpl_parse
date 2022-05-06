package pdf

import (
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func cfRender(c zpl.Command, p *Pdf) *gopdf.GoPdf {
	return &p.pdf
}
