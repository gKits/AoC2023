package day12

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(in []string, part int) (int, error) {
	switch part {
	case 1:
		return part1(in), nil
	case 2:
		return part2(in), nil
	default:
		return -1, fmt.Errorf("day12 does not have a part %d", part)
	}
}

func part1(in []string) int {
	for _, line := range in {
		springs, groups := formatSprings(line)
		fmt.Println(springs, groups, len(springs))
	}
	return 0
}

func part2(in []string) int {
	return 0
}

func formatSprings(line string) ([]string, []int) {
	line = strings.Trim(line, ".")
	line = strings.ReplaceAll(line, ".", " ")
	var prev string
	for line != prev {
		prev = line
		line = strings.ReplaceAll(line, "  ", " ")
	}

	data := strings.Split(line, " ")
	groups := strings.Split(data[len(data)-1], ",")
	nums := make([]int, len(groups))
	for i, n := range groups {
		nums[i], _ = strconv.Atoi(n)
	}
	return data[:len(data)-1], nums
}
