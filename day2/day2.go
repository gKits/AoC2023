package day2

import (
	"fmt"
	"strconv"
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
		return -1, fmt.Errorf("day2 does not have a part %d", part)
	}
}

const (
	red   = 12
	green = 13
	blue  = 14
)

func part1(in []string) int {
	sum := 0
	ch := make(chan int, len(in))
	defer close(ch)

	for _, line := range in {
		go func(s string) {
			ch <- gameIsPossible(s)
		}(line)
	}

	for i := 1; i <= len(in); i++ {
		sum += <-ch
	}

	return sum
}

func gameIsPossible(line string) int {
	scan := scanner.Scanner{}
	scan.Init(strings.NewReader(line))
	scan.Mode = scanner.ScanIdents | scanner.ScanInts

	var game int
	var setGame = false
	var val int
	var err error

	for r := scan.Scan(); r != scanner.EOF; r = scan.Scan() {
		switch r {
		case scanner.Int:
			val, err = strconv.Atoi(scan.TokenText())
			if err != nil {
				panic(err)
			}
			if setGame {
				game = val
				setGame = false
			}
			break
		case scanner.Ident:
			switch scan.TokenText() {
			case "Game":
				setGame = true
			case "red":
				if val > red {
					return 0
				}
				break
			case "blue":
				if val > blue {
					return 0
				}
				break
			case "green":
				if val > green {
					return 0
				}
				break
			}
		}
	}
	return game
}

func part2(in []string) int {
	sum := 0
	ch := make(chan int, len(in))
	defer close(ch)

	for _, line := range in {
		go func(s string) {
			ch <- powerOfGame(s)
		}(line)
	}

	for i := 1; i <= len(in); i++ {
		sum += <-ch
	}

	return sum
}

func powerOfGame(line string) int {
	scan := scanner.Scanner{}
	scan.Init(strings.NewReader(line))
	scan.Mode = scanner.ScanIdents | scanner.ScanInts

	var red int
	var blue int
	var green int
	var val int
	var err error

	for r := scan.Scan(); r != scanner.EOF; r = scan.Scan() {
		switch r {
		case scanner.Int:
			val, err = strconv.Atoi(scan.TokenText())
			if err != nil {
				panic(err)
			}
			break
		case scanner.Ident:
			switch scan.TokenText() {
			case "red":
				if val > red {
					red = val
				}
				break
			case "blue":
				if val > blue {
					blue = val
				}
				break
			case "green":
				if val > green {
					green = val
				}
				break
			}
		}
	}
	return green * blue * red
}
