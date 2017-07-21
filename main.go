package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	var size int
	var out string
	flag.IntVar(&size, "size", 6, "size of the table (number of rows and cols)")
	flag.StringVar(&out, "out", "", "filename for the output, in pdf format")
	flag.Parse()

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)

	data := []string{}
	for i := 1; i <= 120; i++ {
		data = append(data, fmt.Sprintf("entry %d", i))
	}
	pages := randomizeData(data, size*size)

	for _, p := range pages {
		err := generatePage(pdf, p)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := pdf.OutputFileAndClose(out)
	if err != nil {
		log.Fatal(err)
	}
}
