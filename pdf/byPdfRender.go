package pdf

import (
	"log"
	"regexp"
	"strconv"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

func byRender(c zpl.Command, p *Pdf) *gopdf.GoPdf {
	params := c.GetParameters()

	nonStrRE := regexp.MustCompile(`\D*`)

	width, err := strconv.ParseFloat(nonStrRE.ReplaceAllString(params[0], ""), 64)
	if nil != err {
		log.Fatal(err)
	}

	p.barcodeSettings.Width = width

	widthRatio, err := strconv.ParseFloat(nonStrRE.ReplaceAllString(params[1], ""), 64)
	if nil != err {
		log.Fatal(err)
	}

	p.barcodeSettings.WidthRatio = widthRatio

	height, err := strconv.ParseFloat(nonStrRE.ReplaceAllString(params[2], ""), 64)
	if nil != err {
		log.Fatal(err)
	}

	p.barcodeSettings.Height = height

	return &p.pdf
}
