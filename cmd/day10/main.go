package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type op string

const (
	NOOP op = "noop"
	ADDX op = "addx"
)

type operation struct {
	name           op
	value          int
	executedCycles int
}

func (o *operation) execute() {
	o.executedCycles++
}

func (o *operation) isDone() bool {
	switch o.name {
	case NOOP:
		return o.executedCycles == 1
	case ADDX:
		return o.executedCycles == 2
	default:
		return false
	}
}

func part1(operations []operation) int {
	x, cycle, sum := 1, 1, 0
	for _, operation := range operations {
		for !operation.isDone() {
			sum += signalStrength(x, cycle)
			operation.execute()
			cycle++
		}

		x += operation.value
	}

	return sum
}

func signalStrength(x, cycle int) int {
	switch cycle {
	case 20:
		return x * cycle
	case 60:
		return x * cycle
	case 100:
		return x * cycle
	case 140:
		return x * cycle
	case 180:
		return x * cycle
	case 220:
		return x * cycle
	default:
		return 0
	}
}

func readInput() ([]operation, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	ops := make([]operation, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		var v int
		if len(split) > 1 {
			v, _ = strconv.Atoi(split[1])
		}
		ops = append(ops, operation{
			name:  op(split[0]),
			value: v,
		})
	}

	return ops, nil
}
