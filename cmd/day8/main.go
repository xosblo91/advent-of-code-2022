package main

import (
	"bufio"
	"os"
	"strconv"
)

func part1(trees [][]int) int {
	visibleTrees := 0

	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			if visible := isEdge(x, y, trees); visible {
				visibleTrees++
				continue
			}

			if visible := visibleHorizontally(trees[y][x], x, trees[y]); visible {
				visibleTrees++
				continue
			}

			if visible := visibleVertically(trees[y][x], x, y, trees); visible {
				visibleTrees++
				continue
			}
		}
	}

	return visibleTrees
}

func visibleHorizontally(currentPosHeight, xPos int, xRow []int) bool {
	visible := true
	for x, height := range xRow {
		if x == xPos && visible {
			return true
		} else if x == xPos {
			visible = true
		}
		if height >= currentPosHeight && x < xPos {
			visible = false
		}
		if height >= currentPosHeight && x > xPos {
			visible = false
		}
	}

	return visible
}

func visibleVertically(currentPosHeight, xPos, yPos int, trees [][]int) bool {
	visible := true
	for y := 0; y < len(trees); y++ {
		if y == yPos && visible {
			return true
		} else if y == yPos {
			visible = true
		}

		height := trees[y][xPos]
		if height >= currentPosHeight && y < yPos {
			visible = false
		}
		if height >= currentPosHeight && y > yPos {
			visible = false
		}
	}

	return visible
}

func part2(trees [][]int) int {
	scenicScore := 0

	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			if isEdge(x, y, trees) {
				continue
			}

			left, right := visibleHorizontalTrees(trees[y][x], x, trees[y])

			top, bottom := visibleVerticalTrees(trees[y][x], x, y, trees)

			test := left * right * top * bottom
			if test > scenicScore {
				scenicScore = test
			}
		}
	}

	return scenicScore
}

func isEdge(x, y int, trees [][]int) bool {
	return y == 0 || y == len(trees)-1 || x == 0 || x == len(trees[y])-1
}

func visibleHorizontalTrees(currentPosHeight, xPos int, xRow []int) (int, int) {
	visibleTreesLeft := 0
	for i := xPos - 1; i >= 0; i-- {
		if xRow[i] >= currentPosHeight {
			visibleTreesLeft++
			break
		} else {
			visibleTreesLeft++
		}
	}

	visibleTreesRight := 0
	for i := xPos + 1; i < len(xRow); i++ {
		if xRow[i] >= currentPosHeight {
			visibleTreesRight++
			break
		} else {
			visibleTreesRight++
		}
	}

	return visibleTreesLeft, visibleTreesRight
}

func visibleVerticalTrees(currentPosHeight, xPos, yPos int, trees [][]int) (int, int) {
	visibleTop := 0
	for y := yPos - 1; y >= 0; y-- {
		height := trees[y][xPos]
		if height >= currentPosHeight {
			visibleTop++
			break
		} else {
			visibleTop++
		}
	}

	visibleBottom := 0
	for y := yPos + 1; y < len(trees); y++ {
		height := trees[y][xPos]
		if height >= currentPosHeight {
			visibleBottom++
			break
		} else {
			visibleBottom++
		}
	}

	return visibleTop, visibleBottom
}

func readInput() ([][]int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	trees := make([][]int, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := scanner.Text()
		s := make([]int, 0)
		for _, ch := range input {
			n, _ := strconv.Atoi(string(ch))
			s = append(s, n)
		}

		trees = append(trees, s)
	}

	return trees, nil
}
