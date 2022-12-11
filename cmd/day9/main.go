package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	leftDirection  direction = "L"
	rightDirection direction = "R"
	upDirection    direction = "U"
	downDirection  direction = "D"
)

type movement struct {
	direction direction
	distance  int
}

type position struct {
	x int
	y int
}

const (
	part1Knots = 2
	part2Knots = 10
)

func part1(movements []movement) int {
	return simulate(part1Knots, movements)
}

func part2(movements []movement) int {
	return simulate(part2Knots, movements)
}

func simulate(knots int, movements []movement) int {
	visitedTailPositions := make(map[string]bool)
	rope := createRope(knots)

	for _, m := range movements {
		rope, visitedTailPositions = moveRope(rope, m, visitedTailPositions)
	}

	return len(visitedTailPositions)
}

func createRope(knots int) []*position {
	rope := make([]*position, 0, knots)
	for i := 0; i < knots; i++ {
		rope = append(rope, &position{})
	}

	return rope
}

func moveRope(knots []*position, m movement, visitedPositions map[string]bool) ([]*position, map[string]bool) {
	for i := 0; i < m.distance; i++ {
		for i, knot := range knots {
			if i == 0 {
				moveHead(knot, m)
				continue
			}

			moveKnot(knots[i-1], knot)

			if i == len(knots)-1 {
				visitedPositions[fmt.Sprintf("{%d,%d}", knot.x, knot.y)] = true
			}
		}
	}

	return knots, visitedPositions
}

func moveHead(headPos *position, m movement) {
	if m.direction == leftDirection {
		headPos.x -= 1
	} else if m.direction == rightDirection {
		headPos.x += 1
	} else if m.direction == upDirection {
		headPos.y += 1
	} else if m.direction == downDirection {
		headPos.y -= 1
	}
}

func moveKnot(knot1 *position, knot2 *position) {
	xDistance := getDistance(knot2.x, knot1.x)
	yDistance := getDistance(knot2.y, knot1.y)

	if xDistance <= 1 && yDistance <= 1 {
		return
	}
	if knot2.y > knot1.y {
		knot2.y--
	} else if knot2.y < knot1.y {
		knot2.y++
	}
	if knot2.x > knot1.x {
		knot2.x--
	} else if knot2.x < knot1.x {
		knot2.x++
	}
}

func getDistance(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func readInput() ([]movement, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	movements := make([]movement, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		distance, _ := strconv.Atoi(split[1])
		movements = append(movements, movement{
			direction: direction(split[0]),
			distance:  distance,
		})
	}

	return movements, nil
}
