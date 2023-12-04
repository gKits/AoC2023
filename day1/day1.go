package day1

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
		return -1, fmt.Errorf("day1 does not have a part %d", part)
	}
}

func part1(in []string) int {
	sum := 0
	ch := make(chan int, len(in))

	for _, line := range in {
		go func(s string) {
			var f, l int
			for _, c := range s {
				if isDigit(c) {
					if f == 0 {
						f, _ = strconv.Atoi(string(c))
					}
					l, _ = strconv.Atoi(string(c))
				}
			}
			ch <- 10*f + l
		}(line)
	}

	for i := 0; i < len(in); i++ {
		sum += <-ch
	}

	return sum
}

func part2(in []string) int {
	sum := 0
	ch := make(chan int, len(in))

	fwReplace := strings.NewReplacer(
		"one", "one1one",
		"two", "two2two",
		"three", "three3three",
		"four", "four4four",
		"five", "five5five",
		"six", "six6six",
		"seven", "seven7seven",
		"eight", "eight8eight",
		"nine", "nine9nine",
	)

	bwReplace := strings.NewReplacer(
		"eno", "eno1eno",
		"owt", "owt2owt",
		"eerht", "eerht3eerht",
		"ruof", "ruof4ruof",
		"evif", "evif5evif",
		"xis", "xis6xis",
		"neves", "neves7neves",
		"thgie", "thgie8thgie",
		"enin", "enin9enin",
	)

	for _, line := range in {
		fwLine := fwReplace.Replace(line)
		bw := []byte(line)
		slices.Reverse(bw)
		bwLine := bwReplace.Replace(string(bw))

		// fmt.Println(fwLine, bwLine)

		go func(s string) {
			var f int
			for _, c := range fwLine {
				if isDigit(c) {
					f, _ = strconv.Atoi(string(c))
					break
				}
			}

			var l int
			for _, c := range bwLine {
				if isDigit(c) {
					l, _ = strconv.Atoi(string(c))
					break
				}
			}

			ch <- 10*f + l
		}(line)
	}

	for i := 0; i < len(in); i++ {
		sum += <-ch
	}

	return sum
}

func isDigit(r rune) bool {
	return strings.Contains("123456789", string(r))
}
