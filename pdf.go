package main

import "github.com/jung-kurt/gofpdf"

func generatePage(pdf *gofpdf.Fpdf, data []string) error {
	pdf.AddPage()

	for _, v := range data {
		pdf.CellFormat(55, 17, v, "1", 0, "C", false, 0, "")
	}

	return nil
}
