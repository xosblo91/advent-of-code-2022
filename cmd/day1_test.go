package main

import "testing"

func TestPart1(t *testing.T) {
	elves, _ := readInput()
	weight := findTheBiggestChad(elves)
	t.Log(weight)
}

func TestPart2(t *testing.T) {
	elves, _ := readInput()
	weight := findTheThreeBiggestChads(elves)
	t.Log(weight)
}
