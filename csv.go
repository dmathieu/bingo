package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
)

func parseCsv(path string) ([]string, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(c))
	var data []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, record[0])
	}

	return data, nil
}
