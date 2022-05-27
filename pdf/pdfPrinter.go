package pdf

import (
	zpl "zplgun/zpl"

	"github.com/signintech/gopdf"
)

var ZplPdfRenders = map[string]RenderFunc{
	"^FD": fdRender,
	"^LH": lhRender,
	//"^FX": zplCfHandler,pdf
	//"^XA": zplCfHandler,
	//"^FS": zplCfHandler,
	//"^GB": zplCfHandler,
	"^FO": foRender,
	"^CF": cfRender,
	"^FP": fpRender,
	"^BY": byRender,
}

type labelHomeSettings struct {
	X float64
	Y float64
}

type barcodeSettings struct {
	Width      float64
	WidthRatio float64
	Height     float64
}

var pdfFileName = "zpl_output.pdf"

type Pdf struct {
	pdf               gopdf.GoPdf
	zplCommands       []zpl.Command
	labelHomeSettings labelHomeSettings
	barcodeSettings   barcodeSettings
}

type RenderFunc func(c zpl.Command, p *Pdf) *gopdf.GoPdf

func CreatePdf(zpl []zpl.Command) Pdf {
	p := gopdf.GoPdf{}
	p.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	p.AddPage()
	p.AddTTFFont("times", "./times.ttf")
	p.SetFont("times", "", 14)

	return Pdf{
		pdf:               p,
		zplCommands:       zpl,
		labelHomeSettings: labelHomeSettings{X: 0, Y: 0},
		barcodeSettings:   barcodeSettings{Width: 0, WidthRatio: 0, Height: 0},
	}
}

func (p *Pdf) Render() {
	for _, c := range p.zplCommands {
		r := findRender(c)
		if nil != r {
			r(c, p)
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
