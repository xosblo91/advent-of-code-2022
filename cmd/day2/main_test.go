package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := readInput()
	require.NoError(t, err)

	score := play(input)
	t.Log(score)
}

func TestPart2(t *testing.T) {
	input, err := readInput()
	require.NoError(t, err)
	
	score := play2(input)
	t.Log(score)
}
