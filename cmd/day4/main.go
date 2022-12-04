package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type cleaningPair struct {
	Elf1 elf
	Elf2 elf
}

type elf struct {
	Lower int
	Upper int
}

func part1(pairs []cleaningPair) int {
	fullOverlap := 0
	for _, p := range pairs {
		if p.Elf1.Lower <= p.Elf2.Lower && p.Elf1.Upper >= p.Elf2.Upper {
			fullOverlap++
		} else if p.Elf2.Lower <= p.Elf1.Lower && p.Elf2.Upper >= p.Elf1.Upper {
			fullOverlap++
		}
	}

	return fullOverlap
}

func part2(pairs []cleaningPair) int {
	overlap := 0
	for _, p := range pairs {
		if p.Elf1.Lower <= p.Elf2.Upper && p.Elf1.Upper >= p.Elf2.Lower {
			overlap++
		} else if p.Elf2.Lower <= p.Elf1.Upper && p.Elf2.Upper >= p.Elf1.Lower {
			overlap++
		}
	}

	return overlap
}

func parse(input string) cleaningPair {
	split := strings.FieldsFunc(input, func(c rune) bool {
		return c == ',' || c == '-'
	})

	elf1Lower, _ := strconv.Atoi(split[0])
	elf1Upper, _ := strconv.Atoi(split[1])
	elf2Lower, _ := strconv.Atoi(split[2])
	elf2Upper, _ := strconv.Atoi(split[3])

	return cleaningPair{
		Elf1: elf{
			Lower: elf1Lower,
			Upper: elf1Upper,
		},
		Elf2: elf{
			Lower: elf2Lower,
			Upper: elf2Upper,
		},
	}
}

func readInput() ([]cleaningPair, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	pairs := make([]cleaningPair, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pairs = append(pairs, parse(scanner.Text()))
	}

	return pairs, nil
}
