package pdf

import (
	"strconv"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func foRender(c zpl.Command, p *Pdf) *gopdf.GoPdf {
	params := c.GetParameters()

	if len(params) < 2 {
		return &p.pdf
	}

	x, err := strconv.ParseFloat(params[0], 64)

	if nil != err {
		return &p.pdf
	}

	y, err := strconv.ParseFloat(params[1], 64)

	if nil != err {
		return &p.pdf
	}

	labelX, errLhX := strconv.ParseFloat(p.globalSettings["labelHomeX"], 64)
	labelY, errLhY := strconv.ParseFloat(p.globalSettings["labelHomeY"], 64)

	if nil != errLhX || nil != errLhY {
		return &p.pdf
	}

	p.pdf.SetX(labelX + x)
	p.pdf.SetY(labelY + y)

	return &p.pdf
}
