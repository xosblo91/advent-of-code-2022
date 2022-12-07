package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := readInput()
	require.NoError(t, err)

	n, _ := createNodeTree(input)

	result := part1(n)
	t.Log(result)
}

func TestPart2(t *testing.T) {
	input, err := readInput()
	require.NoError(t, err)

	n, _ := createNodeTree(input)

	freeSpace := totalDiscSpace - n.Size
	neededSpace := updateSize - freeSpace
	result := part2(n, neededSpace)
	t.Log(result)
}
