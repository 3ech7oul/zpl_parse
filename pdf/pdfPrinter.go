package pdf

import (
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

var ZplPdfHandlers = map[string]HandlerFunc{
	"^FD": fdRender,
	//"^FX": zplCfHandler,pdf
	//"^XA": zplCfHandler,
	//"^FS": zplCfHandler,
	//"^GB": zplCfHandler,
	"^FO": foRender,
	"^CF": cfRender,
}

type HandlerFunc func(c zpl.Command, p *gopdf.GoPdf) *gopdf.GoPdf

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
