package day3

import (
	"fmt"
	"slices"
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
		return -1, fmt.Errorf("day3 does not have a part %d", part)
	}
}

func part1(in []string) int {
	sum := 0
	var num []byte
	var isPart = false

	for ln := 0; ln < len(in); ln++ {
		for col := 0; col < len(in[ln]); col++ {
			if isDigit(in[ln][col]) {
				num = append(num, in[ln][col])
				if !isPart {
					if hasAdjacentSymbol(in, ln, col) {
						isPart = true
					}
				}
			} else {
				if isPart {
					val, _ := strconv.Atoi(string(num))
					sum += val
				}

				num = []byte{}
				isPart = false

			}
		}
		if isPart {
			val, _ := strconv.Atoi(string(num))
			sum += val
		}

		num = []byte{}
		isPart = false

	}
	return sum
}

func part2(in []string) int {
	var sum int
	for ln := 0; ln < len(in); ln++ {
		for col := 0; col < len(in[ln]); col++ {
			if in[ln][col] == '*' && numAdjacentParts(in, ln, col) == 2 {
				ratio := getGearRatio(in, ln, col)
				sum += ratio
			}
		}
	}
	return sum
}

func hasAdjacentSymbol(in []string, ln, col int) bool {
	adj := getAdjacent(in, ln, col)
	return strings.ContainsAny(adj, "*@!#$%^&/?|=-_+")
}

func getAdjacent(in []string, ln, col int) string {
	xDirs := []int{0}
	yDirs := []int{0}

	if col > 0 {
		xDirs = append(xDirs, -1)
	}

	if col < len(in[0])-1 {
		xDirs = append(xDirs, 1)
	}

	if ln > 0 {
		yDirs = append(yDirs, -1)
	}

	if ln < len(in)-1 {
		yDirs = append(yDirs, 1)
	}

	adj := []byte{}
	for _, y := range yDirs {
		for _, x := range xDirs {
			adj = append(adj, in[ln+y][col+x])
		}
	}

	return string(adj)
}

func isValidSymbol(b byte) bool {
	return b != '.' && !isDigit(b)
}

func isDigit(b byte) bool {
	return strings.Contains("0123456789", string(b))
}

func numAdjacentParts(in []string, ln, col int) int {
	var num int
	lns := len(in)
	cols := len(in[0])

	if col > 0 && isDigit(in[ln][col-1]) {
		num++
	}
	if col < cols-1 && isDigit(in[ln][col+1]) {
		num++
	}

	if ln > 0 {
		if isDigit(in[ln-1][col]) {
			num++
		} else {
			if col > 0 && isDigit(in[ln-1][col-1]) {
				num++
			}
			if col < cols-1 && isDigit(in[ln-1][col+1]) {
				num++
			}
		}
	}
	if ln < lns-1 {
		if isDigit(in[ln+1][col]) {
			num++
		} else {
			if col > 0 && isDigit(in[ln+1][col-1]) {
				num++
			}
			if col < cols-1 && isDigit(in[ln+1][col+1]) {
				num++
			}
		}
	}
	return num
}

func getGearRatio(in []string, ln, col int) int {
	ratio := 1
	lns := len(in)
	cols := len(in[0])

	if col > 0 && isDigit(in[ln][col-1]) {
		ratio *= getNumber(in[ln], col-1)
	}
	if col < cols-1 && isDigit(in[ln][col+1]) {
		ratio *= getNumber(in[ln], col+1)
	}

	if ln > 0 {
		if isDigit(in[ln-1][col]) {
			ratio *= getNumber(in[ln-1], col)
		} else {
			if col > 0 && isDigit(in[ln-1][col-1]) {
				ratio *= getNumber(in[ln-1], col-1)
			}
			if col < cols-1 && isDigit(in[ln-1][col+1]) {
				ratio *= getNumber(in[ln-1], col+1)
			}
		}
	}
	if ln < lns-1 {
		if isDigit(in[ln+1][col]) {
			ratio *= getNumber(in[ln+1], col)
		} else {
			if col > 0 && isDigit(in[ln+1][col-1]) {
				ratio *= getNumber(in[ln+1], col-1)
			}
			if col < cols-1 && isDigit(in[ln+1][col+1]) {
				ratio *= getNumber(in[ln+1], col+1)
			}
		}
	}

	return ratio
}

func getNumber(line string, col int) int {
	b := []byte{}
	for i := col - 1; i >= 0; i-- {
		if !isDigit(line[i]) {
			break
		}
		b = append(b, line[i])
	}
	slices.Reverse(b)

	b = append(b, line[col])

	for i := col + 1; i < len(line); i++ {
		if !isDigit(line[i]) {
			break
		}
		b = append(b, line[i])
	}

	num, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}

	return num
}
