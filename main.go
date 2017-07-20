package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)

	for i := 1; i <= 10; i++ {
		err := generatePage(pdf, []string{"someone", "someone"})
		if err != nil {
			log.Fatal(err)
		}
	}

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
