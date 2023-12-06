package day6

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
		return -1, fmt.Errorf("day6 does not have a part %d", part)
	}
}

func part1(in []string) int {
	timeScan := scanner.Scanner{}
	recordScan := scanner.Scanner{}
	timeScan.Init(strings.NewReader(in[0]))
	recordScan.Init(strings.NewReader(in[1]))

	res := 1

	for ttok, rtok := timeScan.Scan(), recordScan.Scan(); ttok != scanner.EOF && rtok != scanner.EOF; ttok, rtok = timeScan.Scan(), recordScan.Scan() {
		var time int
		var record int

		if ttok == scanner.Int && rtok == scanner.Int {
			time, _ = strconv.Atoi(timeScan.TokenText())
			record, _ = strconv.Atoi(recordScan.TokenText())
		} else {
			continue
		}

		var i int
		for i = 1; i <= time/2; i++ {
			if i*(time-i) > record {
				break
			}
		}

		res *= time - (2*i - 1)
	}

	return res
}

func part2(in []string) int {
	timeScan := scanner.Scanner{}
	recordScan := scanner.Scanner{}
	timeScan.Init(strings.NewReader(in[0]))
	recordScan.Init(strings.NewReader(in[1]))

	time, _ := strconv.Atoi(strings.ReplaceAll(in[0], " ", "")[5:])
	record, _ := strconv.Atoi(strings.ReplaceAll(in[1], " ", "")[9:])

	var i int
	for i = 1; i <= time/2; i++ {
		if i*(time-i) > record {
			break
		}
	}

	return time - (2*i - 1)
}
