package pdf

import (
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

var ZplPdfHandlers = map[string]HandlerFunc{
	"^FD": fdHandler,
	//"^FX": zplCfHandler,
	//"^XA": zplCfHandler,
	//"^FS": zplCfHandler,
	//"^GB": zplCfHandler,
	//"^FO": zplCfHandler,
	"^CF": cfHandler,
}

type HandlerFunc func(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf

type Pdf struct {
	Pdf gopdf.GoPdf
}

func CreatePdf() Pdf {
	p := gopdf.GoPdf{}
	p.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	p.AddPage()

	return Pdf{
		Pdf: p,
	}
}

func fdHandler(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf {
	p.Cell(nil, "您好")

	return p
}

func cfHandler(c zpl.Command, p gopdf.GoPdf) gopdf.GoPdf {
	return p
}
