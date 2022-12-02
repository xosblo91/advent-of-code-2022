package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	OpponentRock    = "A"
	OpponentPaper   = "B"
	OpponentScissor = "C"
)

const (
	YouRock    = "X"
	YouPaper   = "Y"
	YouScissor = "Z"
)

type game struct {
	Opponent string
	You      string
}

func (g game) Play() int {
	if g.Opponent == OpponentRock {
		return g.OpponentPlaysRock()
	} else if g.Opponent == OpponentPaper {
		return g.OpponentPlaysPaper()
	} else {
		return g.OpponentPlaysScissor()
	}
}

func (g game) OpponentPlaysRock() int {
	if g.You == YouRock {
		return 4
	}
	if g.You == YouPaper {
		return 8
	}

	return 3
}

func (g game) OpponentPlaysPaper() int {
	if g.You == YouRock {
		return 1
	}
	if g.You == YouPaper {
		return 5
	}

	return 9
}

func (g game) OpponentPlaysScissor() int {
	if g.You == YouRock {
		return 7
	}
	if g.You == YouPaper {
		return 2
	}

	return 6
}

func (g game) Play2() int {
	if g.Opponent == OpponentRock {
		return g.OpponentPlaysRock2()
	} else if g.Opponent == OpponentPaper {
		return g.OpponentPlaysPaper2()
	} else {
		return g.OpponentPlaysScissor2()
	}
}

func (g game) OpponentPlaysRock2() int {
	if g.You == "X" {
		return 3
	}
	if g.You == "Y" {
		return 4
	}

	return 8
}

func (g game) OpponentPlaysPaper2() int {
	if g.You == "X" {
		return 1
	}
	if g.You == "Y" {
		return 5
	}

	return 9
}

func (g game) OpponentPlaysScissor2() int {
	if g.You == "X" {
		return 2
	}
	if g.You == "Y" {
		return 6
	}

	return 7
}

func sumScore(games []game) int {
	sum := 0

	for _, game := range games {
		sum += game.Play()
	}

	return sum
}

func sumScore2(games []game) int {
	sum := 0

	for _, game := range games {
		sum += game.Play2()
	}

	return sum
}

func readInput() ([]game, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	games := make([]game, 0)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		games = append(games, game{
			Opponent: split[0],
			You:      split[1],
		})
	}

	return games, nil
}
