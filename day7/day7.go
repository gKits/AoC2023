package day7

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
		return -1, fmt.Errorf("day7 does not have a part %d", part)
	}
}

var rank = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var rank2 = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func part1(in []string) int {
	slices.SortFunc(in, handSortFunc)

	var total int
	for i, line := range in {
		splt := strings.Split(line, " ")
		num, _ := strconv.Atoi(splt[1])
		total += num * (i + 1)
	}

	return total
}

func part2(in []string) int {
	slices.SortFunc(in, handSortFuncJoker)

	var total int
	for i, line := range in {
		splt := strings.Split(line, " ")
		num, _ := strconv.Atoi(splt[1])
		total += num * (i + 1)
	}

	return total
}

func handLevel(hand string) int {
	h := map[rune]int{}
	m := 0

	for i := 0; i < 5; i++ {
		if val, ok := h[rune(hand[i])]; ok {
			h[rune(hand[i])] += 1
			if val+1 > m {
				m = val + 1
			}
		} else {
			h[rune(hand[i])] = 1
		}
	}

	switch len(h) {
	case 1:
		return 6
	case 2:
		if m == 4 {
			return 5
		}
		return 4
	case 3:
		if m == 3 {
			return 3
		}
		return 2
	case 4:
		return 1
	default:
		return 0
	}
}

func handLevelJoker(hand string) int {
	h := map[rune]int{}
	m := 0
	var best rune

	for i := 0; i < 5; i++ {
		if val, ok := h[rune(hand[i])]; ok {
			h[rune(hand[i])] += 1
			if val+1 > m {
				m = val + 1
				best = rune(hand[i])
			} else if val+1 == m {
				if rank2[best] < rank2[rune(hand[i])] {
					best = rune(hand[i])
				}
			}
		} else {
			h[rune(hand[i])] = 1
			if 1 > m {
				m = 1
				best = rune(hand[i])
			} else if m == 1 {
				if rank2[best] < rank2[rune(hand[i])] {
					best = rune(hand[i])
				}
			}
		}
	}
	newHand := strings.ReplaceAll(hand, "J", string(best))
	return handLevel(newHand)
}

func handSortFunc(a, b string) int {
	lvlA, lvlB := handLevel(a), handLevel(b)
	if lvlA == lvlB {
		for i := 0; i < 5; i++ {
			rankA, rankB := rank[rune(a[i])], rank[rune(b[i])]
			if rankA != rankB {
				return rankA - rankB
			}
		}
		return 0
	}
	return lvlA - lvlB
}

func handSortFuncJoker(a, b string) int {
	lvlA := handLevelJoker(a)
	lvlB := handLevelJoker(b)
	if lvlA == lvlB {
		for i := 0; i < 5; i++ {
			rankA, rankB := rank2[rune(a[i])], rank2[rune(b[i])]
			if rankA != rankB {
				return rankA - rankB
			}
		}
		return 0
	}
	return lvlA - lvlB
}
