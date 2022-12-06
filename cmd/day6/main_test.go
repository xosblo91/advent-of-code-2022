package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	characters, err := readInput()
	require.NoError(t, err)

	result := part1(4, characters)
	t.Log(result)
}

func TestPart2(t *testing.T) {
	characters, err := readInput()
	require.NoError(t, err)

	result := part1(14, characters)
	t.Log(result)
}
