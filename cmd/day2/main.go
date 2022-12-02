package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	win  = 6
	draw = 3
)

var winConditions = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var drawConditions = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var lossConditions = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

type game struct {
	Opponent string
	Me       string
}

func score(move string) int {
	if move == "X" {
		return 1
	} else if move == "Y" {
		return 2
	}

	return 3
}

func play(games []game) int {
	var sum int
	for _, game := range games {
		if winningMove := winConditions[game.Opponent]; winningMove == game.Me {
			sum += win
		} else if drawMove := drawConditions[game.Opponent]; drawMove == game.Me {
			sum += draw
		}

		sum += score(game.Me)
	}

	return sum
}

func play2(games []game) int {
	var sum int
	for _, game := range games {
		if game.Me == "Z" {
			sum += win + score(winConditions[game.Opponent])
		} else if game.Me == "Y" {
			sum += draw + score(drawConditions[game.Opponent])
		} else {
			sum += score(lossConditions[game.Opponent])
		}
	}

	return sum
}

func readInput() ([]game, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	games := make([]game, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		games = append(games, game{
			Opponent: split[0],
			Me:       split[1],
		})
	}

	return games, nil
}
