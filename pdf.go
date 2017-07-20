package main

import (
	"errors"

	"github.com/jung-kurt/gofpdf"
)

func generatePage(pdf *gofpdf.Fpdf, data []string) error {
	pdf.AddPage()
	n := nbColLines(data)
	if n == 0 {
		return errors.New("Invalid number of data. Cannot calculate equal number of columns and lines")
	}

	for i, v := range data {
		pdf.CellFormat(float64(275/n), float64(179/n), v, "1", 0, "C", false, 0, "")

		if (i+1)%n == 0 {
			pdf.Ln(-1)
		}
	}

	return nil
}

func nbColLines(data []string) int {
	for i := 2; i <= 10; i++ {
		if len(data)/i == i {
			return i
		}
	}
	return 0
}
