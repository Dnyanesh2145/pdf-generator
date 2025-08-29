package handler

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "fmt"
	"io"
	"net/http"
	"pdf-generator-service/services/pdf"
	"sync"

	"github.com/gin-gonic/gin"
)

type Contollers struct {
	pdfService pdf.PdfService
}

type PdfRequest struct {
	Data map[string]any `json:"data" form:"data"`
}

func (ctr *Contollers) GeneratePDFHandler(c *gin.Context) {
	dataJson := c.PostForm("data")
	var request map[string]any
	if err := json.Unmarshal([]byte(dataJson), &request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON in 'data' field"})
		return
	}

	file, _, fileError := c.Request.FormFile("html_file")
	if fileError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HTML file is required"})
		return
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	html := buf.String()
	var (
		pdfBytes []byte
		tmpBytes []byte
		err      error
		wg       sync.WaitGroup
	)
	// fmt.Println("html", len(html),html)
	fmt.Println("data", request)
	// var data map[string]any
	// json.Unmarshal([]byte(dataJson))
	wg.Add(1)
	go func() {
		defer wg.Done()
		tmpBytes, err = ctr.pdfService.RenderTemplate(html, request)
		if err != nil {
			return
		}
		pdfBytes, err = ctr.pdfService.GeneratoratePdf(tmpBytes)
	}()
	wg.Wait()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=generated.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
