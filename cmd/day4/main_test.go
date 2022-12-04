package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	pairs, err := readInput()
	require.NoError(t, err)

	count := part1(pairs)
	t.Log(count)
}

func TestPart2(t *testing.T) {
	pairs, err := readInput()
	require.NoError(t, err)

	count := part2(pairs)
	t.Log(count)
}
