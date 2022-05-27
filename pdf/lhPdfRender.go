package pdf

import (
	"log"
	"regexp"
	"strconv"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func lhRender(c zpl.Command, p *Pdf) *gopdf.GoPdf {
	params := c.GetParameters()

	nonStrRE := regexp.MustCompile(`\D*`)

	x, err := strconv.ParseFloat(nonStrRE.ReplaceAllString(params[0], ""), 64)
	if nil != err {
		log.Fatal(err)
	}

	p.labelHomeSettings.X = x

	y, err := strconv.ParseFloat(nonStrRE.ReplaceAllString(params[1], ""), 64)
	if nil != err {
		log.Fatal(err)
	}

	p.labelHomeSettings.Y = y

	return &p.pdf
}
