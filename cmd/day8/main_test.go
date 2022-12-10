package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	trees, err := readInput()
	require.NoError(t, err)

	result := part1(trees)
	t.Log(result)
}

func TestPart2(t *testing.T) {
	trees, err := readInput()
	require.NoError(t, err)

	result := part2(trees)
	t.Log(result)
}
