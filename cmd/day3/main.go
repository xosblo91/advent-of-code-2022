package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

type rucksack struct {
	Complete     string
	Compartment1 string
	Compartment2 string
}

func part1(rucksacks []rucksack) int {
	sum := 0
	for _, rucksack := range rucksacks {
		found := make(map[rune]bool)
		for _, item := range rucksack.Compartment1 {
			_, exist := found[item]
			if strings.ContainsRune(rucksack.Compartment2, item) && !exist {
				found[item] = true
				sum += score(item)
			}
		}
	}

	return sum
}

func part2(rucksacks []rucksack) int {
	sum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		for _, item := range rucksacks[i].Complete {
			if strings.ContainsRune(rucksacks[i+1].Complete, item) && strings.ContainsRune(rucksacks[i+2].Complete, item) {
				sum += score(item)
				break
			}
		}
	}

	return sum
}

func score(r rune) int {
	if unicode.IsLower(r) {
		return int(1 + r - 'a')
	}

	return int(27 + r - 'A')
}

func readInput() ([]rucksack, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	rucksacks := make([]rucksack, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		items := scanner.Text()
		itemRunes := []rune(items)

		rucksacks = append(rucksacks, rucksack{
			Complete:     items,
			Compartment1: string(itemRunes[0 : len(itemRunes)/2]),
			Compartment2: string(itemRunes[len(itemRunes)/2:]),
		})
	}

	return rucksacks, nil
}
