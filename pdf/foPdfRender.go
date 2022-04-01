package pdf

import (
	"strconv"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func foRenderer(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf {
	params := c.GetParameters()

	if len(params) < 2 {
		return p
	}

	x, err := strconv.ParseFloat(params[0].Value, 64)

	if nil != err {
		return p
	}

	y, err := strconv.ParseFloat(params[1].Value, 64)

	if nil != err {
		return p
	}

	p.SetX(x)
	p.SetY(y)

	return p
}
