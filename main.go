package main

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)

	data := []string{}
	for i := 1; i <= 120; i++ {
		data = append(data, fmt.Sprintf("entry %d", i))
	}
	pages := randomizeData(data, 25)

	for _, p := range pages {
		err := generatePage(pdf, p)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
