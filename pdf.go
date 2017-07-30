package main

import (
	"errors"
	"math"

	"github.com/jung-kurt/gofpdf"
)

func generatePage(pdf *gofpdf.Fpdf, data []string) error {
	pdf.AddPage()
	n := math.Sqrt(float64(len(data)))
	if n != float64(int(n)) {
		return errors.New("Invalid data number. Need a square number to get equal number of columns and lines")
	}

	tr := pdf.UnicodeTranslatorFromDescriptor("")

	for i, v := range data {
		pdf.CellFormat(275/n, 179/n, tr(v), "1", 0, "C", false, 0, "")

		if (i+1)%int(n) == 0 {
			pdf.Ln(-1)
		}
	}

	return nil
}
