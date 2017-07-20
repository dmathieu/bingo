package main

import (
	"log"
	"math/rand"
	"time"
)

func randomizeData(data []string, nb int) [][]string {
	rand.Seed(time.Now().UTC().UnixNano())
	l := len(data)
	nbPages := calcNbPages(l, nb)
	nbOccurences := (nb * nbPages) / l

	result := make([][]string, nbPages)
	for i := 0; i < len(result); i++ {
		result[i] = make([]string, nb)
	}

	for i := 1; i <= nbOccurences; i++ {
		for j := 0; j < len(data); j++ {
			result = injectRecord(result, data[j], 100)
		}
	}

	return result
}

func calcNbPages(l int, nb int) int {
	n := 1

	for {
		if (nb*n)%l == 0 {
			return n
		}
		n++
	}
}

func injectRecord(result [][]string, entry string, m int) [][]string {
	if m == 0 {
		log.Fatal("Too many retries")
	}

	page := rand.Intn(len(result))
	position := rand.Intn(len(result[page]))

	if result[page][position] != "" {
		return injectRecord(result, entry, m)
	}

	for i := 0; i < len(result[page]); i++ {
		if result[page][i] == entry {
			return injectRecord(result, entry, m-1)
		}
	}

	result[page][position] = entry
	return result
}
