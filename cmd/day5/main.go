package main

import (
	"bufio"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"os"
)

//var testStacks = []Stack[string]{
//	NewStack([]string{"Z", "N"}),
//	NewStack([]string{"M", "C", "D"}),
//	NewStack([]string{"P"}),
//}

var stacks = []Stack[string]{
	NewStack([]string{"Q", "S", "W", "C", "Z", "V", "F", "T"}),
	NewStack([]string{"Q", "R", "B"}),
	NewStack([]string{"B", "Z", "T", "Q", "P", "M", "S"}),
	NewStack([]string{"D", "V", "F", "R", "Q", "H"}),
	NewStack([]string{"J", "G", "L", "D", "B", "S", "T", "P"}),
	NewStack([]string{"W", "R", "T", "Z"}),
	NewStack([]string{"H", "Q", "M", "N", "S", "F", "R", "J"}),
	NewStack([]string{"R", "N", "F", "H", "W"}),
	NewStack([]string{"J", "Z", "T", "Q", "P", "R", "B"}),
}

type move struct {
	Amount int
	From   int
	To     int
}

func part1(moves []move) string {
	for _, move := range moves {
		for i := 0; i < move.Amount; i++ {
			s := stacks[move.From-1].Pop()
			stacks[move.To-1].Push(s)
		}
	}

	var result string
	for _, stack := range stacks {
		result += stack.Pop()
	}

	return result
}

func part2(moves []move) string {
	for _, move := range moves {
		crates := make([]string, 0, move.Amount)
		for i := 0; i < move.Amount; i++ {
			crates = append(crates, stacks[move.From-1].Pop())
		}

		crates = pie.Reverse(crates)
		for _, crate := range crates {
			stacks[move.To-1].Push(crate)
		}
	}

	var result string
	for _, stack := range stacks {
		result += stack.Pop()
	}

	return result
}

func readInput() ([]move, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	moves := make([]move, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var amount, from, to int
		_, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &amount, &from, &to)
		if err != nil {
			return nil, err
		}

		moves = append(moves, move{
			Amount: amount,
			From:   from,
			To:     to,
		})
	}

	return moves, nil
}
