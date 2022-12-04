package main

import (
	"bufio"
	"fmt"
	"os"
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

func parse(input string) (cleaningPair, error) {
	var elf1Lower, elf1Upper, elf2Lower, elf2Upper int
	_, err := fmt.Sscanf(input, "%d-%d,%d-%d", &elf1Lower, &elf1Upper, &elf2Lower, &elf2Upper)
	if err != nil {
		return cleaningPair{}, err
	}

	return cleaningPair{
		Elf1: elf{
			Lower: elf1Lower,
			Upper: elf1Upper,
		},
		Elf2: elf{
			Lower: elf2Lower,
			Upper: elf2Upper,
		},
	}, nil
}

func readInput() ([]cleaningPair, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	pairs := make([]cleaningPair, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pair, err := parse(scanner.Text())
		if err != nil {
			return nil, err
		}

		pairs = append(pairs, pair)
	}

	return pairs, nil
}
