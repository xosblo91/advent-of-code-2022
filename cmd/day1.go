package main

import (
	"bufio"
	"os"
	"strconv"
)

func findTheBiggestChad(elves map[int]int) int {
	biggestChadCarryWeight := 0
	biggestChadIndex := 0
	for i, calories := range elves {
		if calories > biggestChadCarryWeight {
			biggestChadCarryWeight = calories
			biggestChadIndex = i
		}
	}

	delete(elves, biggestChadIndex)
	return biggestChadCarryWeight
}

func findTheThreeBiggestChads(elves map[int]int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += findTheBiggestChad(elves)
	}

	return sum
}

func readInput() (map[int]int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	elves := make(map[int]int)

	counter := 0
	sum := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves[counter] = sum
			sum = 0
		} else {
			calories, _ := strconv.Atoi(scanner.Text())
			sum += calories
		}
		counter++
	}

	return elves, nil
}
