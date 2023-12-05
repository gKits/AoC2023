package main

import (
	"AoC2023/day1"
	"AoC2023/day2"
	"AoC2023/day3"
	"AoC2023/day4"
	"AoC2023/day5"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	var day int
	var part int
	var test bool
	flag.IntVar(&day, "day", 0, "input the daily puzzle you want to solve")
	flag.IntVar(&part, "part", 1, "input the part you want to solve")
	flag.BoolVar(&test, "test", false, "use test input instead")
	flag.Parse()

	path := fmt.Sprintf("./inputs/day%d", day)
	if test {
		path = fmt.Sprintf("./inputs/test_day%d", day)
	}
	in, err := loadInput(path)
	if err != nil {
		log.Fatal(err)
	}

	var res int
	switch day {
	case 1:
		res, err = day1.Solve(in, part)
	case 2:
		res, err = day2.Solve(in, part)
	case 3:
		res, err = day3.Solve(in, part)
	case 4:
		res, err = day4.Solve(in, part)
	case 5:
		res, err = day5.Solve(in, part)
	default:
		log.Fatal(fmt.Errorf("invalid day %d/n", day))
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The result of day %d part %d is: %d\nProcess took %.5fms", day, part, res, float64(time.Since(start).Microseconds())/1000)
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
