package day10

import (
	"fmt"
	"slices"
	"strings"
)

func Solve(in []string, part int) (int, error) {
	switch part {
	case 1:
		return part1(in), nil
	case 2:
		return part2(in), nil
	default:
		return -1, fmt.Errorf("day10 does not have a part %d", part)
	}
}

type direction uint8

const (
	north direction = 8 // 1000
	east  direction = 4 // 0100
	south direction = 2 // 0010
	west  direction = 1 // 0001
)

func (dir direction) Reverse() direction {
	return (dir<<2 + dir>>2) & 15
}

var pipes = map[byte]direction{
	'|': 10, // 1010
	'-': 5,  // 0101
	'L': 12, // 1100
	'J': 9,  // 1001
	'7': 3,  // 0011
	'F': 6,  // 0110
}

type coord struct {
	X, Y int
}

func (c *coord) Equals(o coord) bool { return c.X == o.X && c.Y == o.Y }

func (c *coord) InBounds(w, h int) bool { return c.X >= 0 && c.X < w && c.Y >= 0 && c.Y < h }

func (c coord) Move(dir direction) coord {
	switch dir {
	case north:
		c.Y -= 1
		break
	case east:
		c.X += 1
		break
	case south:
		c.Y += 1
		break
	case west:
		c.X -= 1
		break
	}
	return c
}

func part1(in []string) int {
	var (
		lDir, rDir direction = 0, 0
		l, r       coord     = coord{}, coord{}
		i          int       = 1
	)

	for y, line := range in {
		if x := strings.Index(line, "S"); x != -1 {
			l, r, lDir, rDir = findConnections(in, coord{x, y})
			break
		}
	}

	for i = 1; !l.Equals(r); i++ {
		lPipe := pipes[in[l.Y][l.X]]
		rPipe := pipes[in[r.Y][r.X]]

		l = l.Move(lPipe - lDir)
		r = r.Move(rPipe - rDir)

		lDir = (lPipe - lDir).Reverse()
		rDir = (rPipe - rDir).Reverse()
	}

	return i
}

func findConnections(in []string, start coord) (coord, coord, direction, direction) {
	nexts := []coord{}
	dirs := []direction{}
	for dir := west; dir <= north || len(nexts) < 2; dir *= 2 {
		c := start.Move(dir)
		if c.InBounds(len(in[0]), len(in)) {
			if pipe := pipes[in[c.Y][c.X]]; dir.Reverse()&pipe != 0 {
				nexts = append(nexts, c)
				dirs = append(dirs, dir.Reverse())
			}
		}
	}

	return nexts[0], nexts[1], dirs[0], dirs[1]
}

func firstConnection(in []string, start coord) (coord, direction) {
	for dir := west; dir <= north; dir *= 2 {
		c := start.Move(dir)
		if c.InBounds(len(in[0]), len(in)) {
			if pipe := pipes[in[c.Y][c.X]]; dir.Reverse()&pipe != 0 {
				return c, dir.Reverse()
			}
		}
	}
	return start, 0
}

func part2(in []string) int {
	var (
		dir   direction
		c     coord
		start coord
		maxX  int
		maxY  int
		minX  int
		minY  int
		loop  []coord
	)

	for y, line := range in {
		if x := strings.Index(line, "S"); x != -1 {
			start = coord{x, y}
			maxX = start.X
			minX = start.X
			maxY = start.Y
			minY = start.Y
			loop = []coord{start}
			c, dir = firstConnection(in, start)
			break
		}
	}

	for !c.Equals(start) {
		loop = append(loop, c)
		if c.X > maxX {
			maxX = c.X
		} else if c.X < minX {
			minX = c.X
		}

		if c.Y > maxY {
			maxY = c.Y
		} else if c.Y < minY {
			minY = c.Y
		}

		pipe := pipes[in[c.Y][c.X]]
		c = c.Move(pipe - dir)
		dir = (pipe - dir).Reverse()

	}

	// fmt.Println(loop)
	counter := 0

	for y := minY; y <= maxY; y++ {

		nPipes := 0
		for x := minX; x <= maxX; x++ {
			fmt.Println(y, x, nPipes, counter)
			if !loopContains(loop, x, y) {
				if nPipes%2 != 0 {
					counter++
				}
			} else {
				// pipe := pipes[in[y][x]]
				if in[y][x] != '-' {
					nPipes++
				}
			}
		}
	}

	return counter
}

func loopContains(loop []coord, x, y int) bool {
	return slices.ContainsFunc(loop, func(c coord) bool { return c.X == x && c.Y == y })
}

func horizontalConnection(a, b direction) bool {
	return a<<2&b == 4 || a>>2&b == 1
}
