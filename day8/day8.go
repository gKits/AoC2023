package day8

import (
	"fmt"
)

func Solve(in []string, part int) (int, error) {
	switch part {
	case 1:
		return part1(in), nil
	case 2:
		return part2(in), nil
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

	calcSteps := func(node string) int {
		for i := 0; ; i++ {
			if node[2] == 'Z' {
				return i
			}
			dir := rune(in[0][i%len(in[0])])
			n := nodes[node]
			switch dir {
			case 'L':
				node = n.L
			case 'R':
				node = n.R
			}
		}
	}

	fmt.Println(nodes)

	steps := []int{}
	for node := range nodes {
		if node[2] == 'A' {
			steps = append(steps, calcSteps(node))
		}
	}

	fmt.Println(steps)
	fmt.Println(leastCommonMul(2, 3))
	fmt.Println(greatCommonDiv(2, 3))

	return leastCommonMul(steps[0], steps[1:]...)
}

func leastCommonMul(a int, nums ...int) int {
	if len(nums) >= 1 {
		res := a * nums[0] / greatCommonDiv(a, nums[0])
		return leastCommonMul(res, nums[1:]...)
	}
	return a
}

func greatCommonDiv(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
