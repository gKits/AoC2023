package day5

import (
	"fmt"
	"slices"
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
		return -1, fmt.Errorf("day5 does not have a part %d", part)
	}
}

func part1(in []string) int {
	var header string
	seeds := []int{}
	soil := []mapRange{}
	fert := []mapRange{}
	water := []mapRange{}
	light := []mapRange{}
	temp := []mapRange{}
	humid := []mapRange{}
	locat := []mapRange{}

	for _, line := range in {
		scan := scanner.Scanner{}
		scan.Init(strings.NewReader(line))

		for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
			switch tok {
			case scanner.Int:
				switch header {
				case "seeds":
					num, _ := strconv.Atoi(scan.TokenText())
					seeds = append(seeds, num)
				case "soil":
					soil = append(soil, scanMapRange(&scan))
				case "fertilizer":
					fert = append(fert, scanMapRange(&scan))
				case "water":
					water = append(water, scanMapRange(&scan))
				case "light":
					light = append(light, scanMapRange(&scan))
				case "temperature":
					temp = append(temp, scanMapRange(&scan))
				case "humidity":
					humid = append(humid, scanMapRange(&scan))
				case "location":
					locat = append(locat, scanMapRange(&scan))
				}
			case scanner.Ident:
				if scan.TokenText() != "map" {
					header = scan.TokenText()
				}
			}
		}
	}

	slices.SortFunc(soil, sortMapFunc)
	slices.SortFunc(fert, sortMapFunc)
	slices.SortFunc(water, sortMapFunc)
	slices.SortFunc(light, sortMapFunc)
	slices.SortFunc(temp, sortMapFunc)
	slices.SortFunc(humid, sortMapFunc)
	slices.SortFunc(locat, sortMapFunc)

	var minLoc int
	for _, seed := range seeds {
		loc := getMappedVal(locat, getMappedVal(humid, getMappedVal(temp, getMappedVal(light, getMappedVal(water, getMappedVal(fert, getMappedVal(soil, seed)))))))
		if minLoc == 0 || loc < minLoc {
			minLoc = loc
		}
	}

	return minLoc
}

func part2(in []string) int {
	var header string
	seeds := []int{}
	soil := []mapRange{}
	fert := []mapRange{}
	water := []mapRange{}
	light := []mapRange{}
	temp := []mapRange{}
	humid := []mapRange{}
	locat := []mapRange{}

	for _, line := range in {
		scan := scanner.Scanner{}
		scan.Init(strings.NewReader(line))

		for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
			switch tok {
			case scanner.Int:
				switch header {
				case "seeds":
					num, _ := strconv.Atoi(scan.TokenText())
					seeds = append(seeds, num)
				case "soil":
					soil = append(soil, scanMapRange(&scan))
				case "fertilizer":
					fert = append(fert, scanMapRange(&scan))
				case "water":
					water = append(water, scanMapRange(&scan))
				case "light":
					light = append(light, scanMapRange(&scan))
				case "temperature":
					temp = append(temp, scanMapRange(&scan))
				case "humidity":
					humid = append(humid, scanMapRange(&scan))
				case "location":
					locat = append(locat, scanMapRange(&scan))
				}
			case scanner.Ident:
				if scan.TokenText() != "map" {
					header = scan.TokenText()
				}
			}
		}
	}

	slices.SortFunc(soil, sortMapFunc)
	slices.SortFunc(fert, sortMapFunc)
	slices.SortFunc(water, sortMapFunc)
	slices.SortFunc(light, sortMapFunc)
	slices.SortFunc(temp, sortMapFunc)
	slices.SortFunc(humid, sortMapFunc)
	slices.SortFunc(locat, sortMapFunc)

	var minLoc int
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			loc := getMappedVal(locat, getMappedVal(humid, getMappedVal(temp, getMappedVal(light, getMappedVal(water, getMappedVal(fert, getMappedVal(soil, seed)))))))
			if minLoc == 0 || loc < minLoc {
				minLoc = loc
			}
		}
	}

	return minLoc
}

type mapRange struct {
	Dest int
	Src  int
	Rng  int
}

func getMappedVal(m []mapRange, key int) int {
	for _, rg := range m {
		if key < rg.Src {
			return key
		} else if key >= rg.Src && key < rg.Src+rg.Rng {
			return key - rg.Src + rg.Dest
		}
	}
	return key
}

func sortMapFunc(a, b mapRange) int {
	return a.Src - b.Src
}

func scanMapRange(scan *scanner.Scanner) mapRange {
	rg := mapRange{}
	rg.Dest, _ = strconv.Atoi(scan.TokenText())
	scan.Scan()
	rg.Src, _ = strconv.Atoi(scan.TokenText())
	scan.Scan()
	rg.Rng, _ = strconv.Atoi(scan.TokenText())
	return rg
}
