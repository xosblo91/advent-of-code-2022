package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	movements, err := readInput()
	require.NoError(t, err)

	count := part1(movements)
	t.Log(count)
}

func TestPart2(t *testing.T) {
	movements, err := readInput()
	require.NoError(t, err)

	count := part2(movements)
	t.Log(count)
}
