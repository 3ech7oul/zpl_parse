package pdf

import (
	"strconv"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

var ZplPdfHandlers = map[string]HandlerFunc{
	"^FD": fdHandler,
	//"^FX": zplCfHandler,pdf
	//"^XA": zplCfHandler,
	//"^FS": zplCfHandler,
	//"^GB": zplCfHandler,
	"^FO": foHandler,
	"^CF": cfHandler,
}

type HandlerFunc func(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf

type Pdf struct {
	pdf         gopdf.GoPdf
	zplCommands map[int]zpl.Command
}

func CreatePdf(zpl map[int]zpl.Command) Pdf {
	p := gopdf.GoPdf{}
	p.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	p.AddPage()

	return Pdf{
		pdf:         p,
		zplCommands: zpl,
	}
}

func fdHandler(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf {
	params := c.GetParameters()

	if len(params) < 1 {
		return p
	}

	p.Cell(nil, params[0].Value)

	return p
}

func cfHandler(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf {
	return p
}

func foHandler(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf {
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
