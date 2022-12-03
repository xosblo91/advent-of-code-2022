package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	rucksacks, err := readInput()
	require.NoError(t, err)

	sum := part1(rucksacks)
	t.Log(sum)
}

func TestPart2(t *testing.T) {
	rucksacks, err := readInput()
	require.NoError(t, err)

	sum := part2(rucksacks)
	t.Log(sum)
}
