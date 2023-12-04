package day4

import (
	"fmt"
	"math"
	"strings"
	"text/scanner"
)

func Solve(in []string, part int) (int, error) {
	switch part {
	case 1:
		return part1(in), nil
	case 2:
		return part2(in), nil
	default:
		return -1, fmt.Errorf("day4 does not have a part %d", part)
	}
}

func part1(in []string) int {
	var points int

	for _, card := range in {
		scan := scanner.Scanner{}
		scan.Init(strings.NewReader(card))
		scan.Mode = scanner.ScanInts | scanner.ScanChars | scanner.ScanIdents

		winners := map[string]struct{}{}
		wins := 0
		scanning := false
		scanWinners := true

		for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
			switch tok {
			case scanner.Int:
				if scanning && scanWinners {
					winners[scan.TokenText()] = struct{}{}
				} else if scanning {
					if _, ok := winners[scan.TokenText()]; ok {
						wins++
					}
				}
			case '|':
				scanWinners = false
			case ':':
				scanning = true
			}
		}
		points += int(math.Pow(2, float64(wins))) / 2
	}

	return points
}

func part2(in []string) int {
	extra := make([]int, len(in))

	for i, card := range in {
		scan := scanner.Scanner{}
		scan.Init(strings.NewReader(card))
		scan.Mode = scanner.ScanInts | scanner.ScanChars | scanner.ScanIdents

		winners := map[string]struct{}{}
		wins := 0
		scanning := false
		scanWinners := true

		for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
			switch tok {
			case scanner.Int:
				if scanning && scanWinners {
					winners[scan.TokenText()] = struct{}{}
				} else if scanning {
					if _, ok := winners[scan.TokenText()]; ok {
						wins++
					}
				}
			case '|':
				scanWinners = false
			case ':':
				scanning = true
			}
		}
		for j := i + 1; j < len(in) && j <= i+wins; j++ {
			extra[j] += 1 + extra[i]
		}
	}

	res := len(extra)
	for _, e := range extra {
		res += e
	}

	return res
}
