package day9

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
	// case 2:
	// 	return part2(in), nil
	default:
		return -1, fmt.Errorf("day9 does not have a part %d", part)
	}
}

func part1(in []string) int {
	ch := make(chan int, len(in))
	defer close(ch)

	// for _, line := range in {
	// 	seq := formatSequence(line)
	// 	go func(seq []int) {
	// 		ch <- extrapolateSequence(seq)
	// 	}(seq)
	// }
	//
	// var sum int
	// for i := 0; i < len(in); i++ {
	// 	ext := <-ch
	// 	sum += ext
	// }

	var sum int
	for _, line := range in {
		seq := formatSequence(line)
		ext := extrapolateSequence(seq)
		// fmt.Println(append(seq, ext))
		sum += ext
	}

	return sum
}

func formatSequence(line string) []int {
	scan := scanner.Scanner{}
	scan.Init(strings.NewReader(line))
	scan.Mode = scanner.ScanInts
	seq := []int{}

	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		switch tok {
		case scanner.Int:
			num, _ := strconv.Atoi(scan.TokenText())
			seq = append(seq, num)
		}
	}
	return seq

}

func extrapolateSequence(seq []int) int {
	var next int

	cur := seq
	for {
		cur = diffSeq(cur)
		if allZero(cur) {
			break
		}
		next += cur[len(cur)-1]
	}

	return seq[len(seq)-1] + next
}

func allZero(seq []int) bool {
	for _, n := range seq {
		if n != 0 {
			return false
		}
	}
	return true
}

func diffSeq(seq []int) []int {
	diff := make([]int, len(seq)-1)

	for i := 0; i < len(seq)-1; i++ {
		diff[i] = seq[i+1] - seq[i]
	}

	return diff
}
