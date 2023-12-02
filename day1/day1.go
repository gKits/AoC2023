package day1

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Solve(in []string, part int) (int, error) {
	switch part {
	case 1:
		return part1(in), nil
	case 2:
		return part2(in), nil
	default:
		return 0, fmt.Errorf("day1 does not have a part %d", part)
	}
}

func part1(in []string) int {
	sum := 0
	ch := make(chan int, len(in))

	for _, line := range in {
		go func(s string) {
			f, err := firstDig(s)
			if err != nil {
				panic(err)
			}

			l, err := lastDig(s)
			if err != nil {
				panic(err)
			}

			ch <- 10*f + l
		}(line)
	}

	for i := 0; i < len(in); i++ {
		sum += <-ch
	}

	return sum
}

func firstDig(s string) (int, error) {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return strconv.Atoi(string(r))
		}
	}
	return -1, errors.New("no digits found")
}

func lastDig(s string) (int, error) {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return strconv.Atoi(string(s[i]))
		}
	}
	return -1, errors.New("no digits found")
}

func part2(in []string) int {
	sum := 0
	// ch := make(chan int, len(in))

	// for _, line := range in {
	// 	go func(s string) {
	// 		f, err := firstTextDig(s)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	//
	// 		l, err := lastTextDig(s)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	//
	// 		ch <- 10*f + l
	// 	}(line)
	// }
	//
	// for i := 0; i < len(in); i++ {
	// 	sum += <-ch
	// }

	for _, line := range in {
		f, err := firstTextDig(line)
		if err != nil {
			panic(err)
		}

		l, err := lastTextDig(line)
		if err != nil {
			panic(err)
		}

		log.Println(line, f, l)

		sum += 10*f + l
	}

	return sum
}

var textDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func firstTextDig(s string) (int, error) {
	buf := []byte{}
	for _, r := range s {
		if unicode.IsDigit(r) {
			return strconv.Atoi(string(r))
		} else {
			buf = append(buf, byte(r))
			if !digitStartsWith(string(buf)) {
				buf = []byte{}
				continue
			}
			digit, ok := textDigits[string(buf)]
			if ok {
				return digit, nil
			}
		}
	}
	return -1, errors.New("no digits found")
}

func lastTextDig(s string) (int, error) {
	buf := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		if unicode.IsDigit(rune(r)) {
			return strconv.Atoi(string(r))
		} else {
			buf = slices.Insert(buf, 0, r)
			if !digitEndsWith(string(buf)) {
				buf = []byte{}
				continue
			}
			digit, ok := textDigits[string(buf)]
			if ok {
				return digit, nil
			}
		}
	}
	return -1, errors.New("no digits found")
}

func digitStartsWith(s string) bool {
	for digit := range textDigits {
		if strings.HasPrefix(digit, s) {
			return true
		}
	}
	return false
}

func digitEndsWith(s string) bool {
	for digit := range textDigits {
		if strings.HasSuffix(digit, s) {
			return true
		}
	}
	return false
}
