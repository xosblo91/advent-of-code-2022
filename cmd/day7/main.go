package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	totalDiscSpace = 70000000
	updateSize     = 30000000
)

const (
	cdPrefix       = "$ cd"
	navigateUp     = ".."
	navigateToRoot = "/"
	directory      = "dir"
)

type node struct {
	Parent   *node
	Children []*node
	Name     string
	Size     int
}

func part1(n *node) int {
	total := 0

	if n.Size <= 100000 {
		total += n.Size
	}

	for _, file := range n.Children {
		total += part1(file)
	}

	return total
}

func part2(n *node, size int) int {
	nodes := findNodes(n, size)

	smallest := math.MaxInt

	for _, candidate := range nodes {
		if candidate.Size < smallest {
			smallest = candidate.Size
		}
	}

	return smallest
}

func findNodes(n *node, size int) []*node {
	nodes := make([]*node, 0)

	if n.Size >= size {
		nodes = append(nodes, n)
	}

	for _, child := range n.Children {
		nodes = append(nodes, findNodes(child, size)...)
	}

	return nodes
}

func increaseParentSize(node *node, size int) {
	if node.Parent != nil {
		node.Parent.Size += size
		increaseParentSize(node.Parent, size)
	}
}

func changeDirectory(current, root *node, cmd string) *node {
	cmdSplit := strings.Split(cmd, " ")
	cd := cmdSplit[2]

	if cd == navigateToRoot {
		return root
	} else if cd == navigateUp {
		return current.Parent
	} else {
		for _, child := range current.Children {
			if child.Name == cd {
				return child
			}
		}
	}

	return current
}

func list(index int, current *node, line string, lines []string) (*node, int) {
	for {
		if index == len(lines)-1 || strings.Contains(lines[index+1], "$") {
			return current, index
		}

		index++
		line = lines[index]
		lineSplit := strings.Split(line, " ")

		if lineSplit[0] == directory {
			current.Children = append(current.Children, &node{
				Name:   lineSplit[1],
				Parent: current,
			})
		} else {
			size, _ := strconv.Atoi(lineSplit[0])
			current.Size += size
			increaseParentSize(current, size)
		}
	}
}

func createNodeTree(lines []string) (*node, error) {
	root := &node{
		Parent:   nil,
		Children: []*node{},
		Name:     "/",
		Size:     0,
	}

	current := root

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, cdPrefix) {
			current = changeDirectory(current, root, line)
		} else {
			current, i = list(i, current, line, lines)
		}
	}

	return root, nil
}

func readInput() ([]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	lines := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
