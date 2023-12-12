package day11

import (
	"fmt"
	"strings"
)

func Solve(in []string, part int) (int, error) {
	switch part {
	case 1:
		return part1(in), nil
	case 2:
		return part2(in), nil
	default:
		return -1, fmt.Errorf("day11 does not have a part %d", part)
	}
}

type coord struct {
	X, Y int
}

func part1(in []string) int {
	return calculatePathSum(in, 1)
}

func part2(in []string) int {
	return calculatePathSum(in, 1000000-1)
}

func calculatePathSum(in []string, mod int) int {
	var galaxies []coord

	emptyCols := make([]bool, len(in[0]))
	for i := 0; i < len(in[0]); i++ {
		empty := true
		for j := 0; j < len(in); j++ {
			if in[j][i] == '#' {
				empty = false
				break
			}
		}
		emptyCols[i] = empty
	}

	emptyRows := make([]bool, len(in))
	for i, line := range in {
		if !strings.ContainsRune(line, '#') {
			emptyRows[i] = true
		} else {
			sub := []rune(in[i])
			for x := strings.IndexRune(string(sub), '#'); x != -1; x = strings.IndexRune(string(sub), '#') {
				galaxies = append(galaxies, coord{x, i})
				sub[x] = '.'
			}
		}
	}

	var sum int
	for i, galaxyA := range galaxies {
		for _, galaxyB := range galaxies[i:] {
			dx := abs((galaxyA.X + countTo(emptyCols, true, galaxyA.X)*mod) -
				(galaxyB.X + countTo(emptyCols, true, galaxyB.X)*mod))
			dy := abs((galaxyA.Y + countTo(emptyRows, true, galaxyA.Y)*mod) -
				(galaxyB.Y + countTo(emptyRows, true, galaxyB.Y)*mod))
			sum += dx + dy
		}
	}

	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countTo[S []T, T comparable](s S, of T, lim int) int {
	var count int
	for i := 0; i < lim && i < len(s); i++ {
		if s[i] == of {
			count++
		}
	}
	return count
}
