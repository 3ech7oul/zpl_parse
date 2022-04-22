package pdf

import (
	"fmt"
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

var ZplPdfRenders = map[string]RenderFunc{
	"^FD": fdRender,
	//"^FX": zplCfHandler,pdf
	//"^XA": zplCfHandler,
	//"^FS": zplCfHandler,
	//"^GB": zplCfHandler,
	"^FO": foRender,
	"^CF": cfRender,
}

var pdfFileName = "zpl_output.pdf"

type RenderFunc func(c zpl.Command, p *gopdf.GoPdf) *gopdf.GoPdf

type Pdf struct {
	pdf         gopdf.GoPdf
	zplCommands []zpl.Command
}

func CreatePdf(zpl []zpl.Command) Pdf {
	p := gopdf.GoPdf{}
	p.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	p.AddPage()
	p.AddTTFFont("times", "./times.ttf")
	p.SetFont("times", "", 14)

	return Pdf{
		pdf:         p,
		zplCommands: zpl,
	}
}

func (p *Pdf) Render() {
	for _, c := range p.zplCommands {
		fmt.Println(c)
		r := findRender(c)
		if nil != r {
			r(c, &p.pdf)
		}
	}
}

func (p *Pdf) Save() {
	p.pdf.WritePdf(pdfFileName)
}

func findRender(c zpl.Command) RenderFunc {
	result, ok := ZplPdfRenders[c.ZplCommToken]
	if ok {
		return result
	}

	return nil
}
