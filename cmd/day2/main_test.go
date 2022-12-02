package main

import "testing"

func TestPart1(t *testing.T) {
	input, _ := readInput()
	score := sumScore(input)
	t.Log(score)
}

func TestPart2(t *testing.T) {
	input, _ := readInput()
	score := sumScore2(input)
	t.Log(score)
}
