package pdfGenerator

import (
	"bytes"
	"eh_teh_mewa/helperMain"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type RequestPdf struct {
	body string
}

func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

func (pdf *RequestPdf) ParseTempalate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	helperMain.PanicIfError(err)
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	pdf.body = buf.String()
	return nil
}

func (pdf *RequestPdf) GeneratePDF(pdfPath string, args []string) (bool, error) {
	t := time.Now().Unix()
	if _, err := os.Stat("cloneTemplate/"); os.IsNotExist(err) {
		errDir := os.Mkdir("cloneTemplate/", 0777)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}
	err1 := ioutil.WriteFile("cloneTemplate/"+strconv.FormatInt(int64(t), 10)+".html", []byte(pdf.body), 0644)
	if err1 != nil {
		panic(err1)
	}

	f, err := os.Open("cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Use arguments to customize PDF generation process
	for _, arg := range args {
		switch arg {
		case "low-quality":
			pdfg.LowQuality.Set(true)
		case "no-pdf-compression":
			pdfg.NoPdfCompression.Set(true)
		case "grayscale":
			pdfg.Grayscale.Set(true)
			// Add other arguments as needed
		}
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir + "/cloneTemplate")

	return true, nil
}
