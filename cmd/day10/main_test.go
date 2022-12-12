package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	ops, err := readInput()
	require.NoError(t, err)

	result := part1(ops)
	t.Log(result)
}
