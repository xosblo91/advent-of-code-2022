package day1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	elves, err := readInput()
	require.NoError(t, err)

	weight := elves.findTheBiggestChad()
	t.Log(weight)
}

func TestPart2(t *testing.T) {
	elves, err := readInput()
	require.NoError(t, err)

	weight := elves.findTheThreeBiggestChads()
	t.Log(weight)
}
