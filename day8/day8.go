package day8

import "fmt"

func Solve(in []string, part int) (int, error) {
	switch part {
	case 1:
		return part1(in), nil
	default:
		return -1, fmt.Errorf("day8 does not have a part %d", part)
	}
}

type node struct {
	L string
	R string
}

func part1(in []string) int {
	nodes := map[string]node{}
	for i := 2; i < len(in); i++ {
		nodes[in[i][:3]] = node{L: in[i][7:10], R: in[i][12:15]}
	}

	var cur = "AAA"
	var dir rune
	var i int

	for i = 0; cur != "ZZZ"; i++ {
		dir = rune(in[0][i%len(in[0])])
		node := nodes[cur]
		switch dir {
		case 'L':
			cur = node.L
		case 'R':
			cur = node.R
		}
	}

	return i
}

func part2(in []string) int {
	nodes := map[string]node{}
	for i := 2; i < len(in); i++ {
		nodes[in[i][:3]] = node{L: in[i][7:10], R: in[i][12:15]}
	}

	var cur = "AAA"
	var dir rune
	var i int

	for i = 0; cur != "ZZZ"; i++ {
		dir = rune(in[0][i%len(in[0])])
		node := nodes[cur]
		switch dir {
		case 'L':
			cur = node.L
		case 'R':
			cur = node.R
		}
	}

	return i
}
