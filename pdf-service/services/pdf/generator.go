package pdf

import (
	"bytes"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PdfService struct{}

func (pdf *PdfService) GeneratoratePdf(html []byte) ([]byte, error) {
   
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(html))

	pdfg.AddPage(page)

	if err := pdfg.Create(); err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}

func (pdf *PdfService) RenderTemplate(html string, data any) ([]byte, error) {

	tmp, err := template.New("client-template").Parse(html)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmp.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
