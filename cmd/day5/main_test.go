package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	moves, err := readInput()
	require.NoError(t, err)

	result := part1(moves)
	t.Log(result)
}

func TestPart2(t *testing.T) {
	moves, err := readInput()
	require.NoError(t, err)

	result := part2(moves)
	t.Log(result)
}
