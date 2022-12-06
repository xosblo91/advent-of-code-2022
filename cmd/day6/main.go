package main

import (
	"bufio"
	"github.com/elliotchance/pie/v2"
	"os"
)

func part1(n int, characters []byte) int {
	count := 0
	for i := n - 1; i < len(characters); i++ {
		sequence := characters[i-n+1 : n+count]
		if pie.AreUnique(sequence) {
			return i + 1
		}

		count++
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
