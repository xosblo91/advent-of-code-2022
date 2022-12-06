package main

import (
	"bufio"
	"github.com/elliotchance/pie/v2"
	"os"
)

func part1(n int, characters []byte) int {
	for i := range characters {
		if i < n-1 {
			continue
		}

		sequence := characters[i-n+1 : i+1]
		if pie.AreUnique(sequence) {
			return i + 1
		}
	}

	return 0
}

func readInput() ([]byte, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	characters := ""
	for scanner.Scan() {
		characters += scanner.Text()
	}

	return []byte(characters), nil
}
