package main

import (
	"math/rand"
	"time"
)

func randomizeData(data []string, nb int) [][]string {
	rand.Seed(time.Now().Unix())
	l := len(data)
	nbPages := calcNbPages(l, nb)
	nbOccurences := (nb * nbPages) / l

	var expended = []string{}
	for i := 1; i <= nbOccurences; i++ {
		for j := 0; j < len(data); j++ {
			expended = append(expended, data[j])
		}
	}

	result := make([][]string, nbPages)
	for i := 0; i < len(result); i++ {
		result[p] = make([]string, nb)
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
