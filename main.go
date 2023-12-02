package main

import (
	"AoC2023/day1"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var day int
	var part int
	flag.IntVar(&day, "day", 0, "input the daily puzzle you want to solve")
	flag.IntVar(&part, "part", 1, "input the part you want to solve")
	flag.Parse()

	in, err := loadInput(fmt.Sprintf("./inputs/day%d", day))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	switch day {
	case 1:
		res, err := day1.Solve(in, part)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
		break
	default:
		fmt.Printf("invalid day %d/n", day)
	}
}

func loadInput(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)

	scan.Split(bufio.ScanLines)
	var in []string

	for scan.Scan() {
		in = append(in, scan.Text())
	}

	return in, nil
}
