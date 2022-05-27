package pdf

import (
	"log"
	"strconv"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func foRender(c zpl.Command, p *Pdf) *gopdf.GoPdf {
	params := c.GetParameters()

	x, err := strconv.ParseFloat(params[0], 64)

	if nil != err {
		log.Fatal(err)
	}

	y, err := strconv.ParseFloat(params[1], 64)

	if nil != err {
		log.Fatal(err)
	}

	p.pdf.SetX(p.labelHomeSettings.X + x)
	p.pdf.SetY(p.labelHomeSettings.Y + y)

	return &p.pdf
}
