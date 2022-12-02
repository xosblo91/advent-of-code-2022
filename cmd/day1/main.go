package day1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type elves []int

func (e elves) findTheBiggestChad() int {
	biggestChad := 0

	for _, elf := range e {
		if elf > biggestChad {
			biggestChad = elf
		}
	}

	return biggestChad
}

func (e elves) findTheThreeBiggestChads() int {
	sort.Slice(e, func(i, j int) bool {
		return e[i] > e[j]
	})

	return e[0] + e[1] + e[2]
}

func readInput() (elves, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	elves := make(elves, 0)
	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			calories, _ := strconv.Atoi(scanner.Text())
			sum += calories
		}
	}

	return elves, nil
}
