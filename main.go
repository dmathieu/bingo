package main

import (
	"flag"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	var size int
	var in, out string
	flag.IntVar(&size, "size", 5, "size of the table (number of rows and cols)")
	flag.StringVar(&in, "in", "", "input filename, in csv format")
	flag.StringVar(&out, "out", "", "filename for the output, in pdf format")
	flag.Parse()

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 12)

	data, err := parseCsv(in)
	if err != nil {
		log.Fatal(err)
	}
	pages, err := randomizeData(data, size*size)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pages {
		err := generatePage(pdf, p)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = pdf.OutputFileAndClose(out)
	if err != nil {
		log.Fatal(err)
	}
}
