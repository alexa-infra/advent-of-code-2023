package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	n, m := len(lines), len(lines[0])
	type pair struct {
		x, y int
	}
	extendUniverse := func(times int) []pair {
		coords := []pair{}
		for i, line := range lines {
			for j, ch := range line {
				if ch == '#' {
					p := pair{j, i}
					coords = append(coords, p)
				}
			}
		}
		for i := n - 1; i >= 0; i-- {
			empty := true
			for j := 0; j < m; j++ {
				if lines[i][j] == '#' {
					empty = false
					break
				}
			}
			if empty {
				for j, coord := range coords {
					if coord.y >= i {
						coords[j].y += times - 1
					}
				}
			}
		}
		for i := m - 1; i >= 0; i-- {
			empty := true
			for j := 0; j < n; j++ {
				if lines[j][i] == '#' {
					empty = false
					break
				}
			}
			if empty {
				for j, coord := range coords {
					if coord.x >= i {
						coords[j].x += times - 1
					}
				}
			}
		}
		return coords
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	dist := func(a, b pair) int {
		return abs(a.x-b.x) + abs(a.y-b.y)
	}
	sumMinDistances := func(coords []pair) int {
		r := 0
		cn := len(coords)
		for i, a := range coords {
			for j := i + 1; j < cn; j++ {
				b := coords[j]
				d := dist(a, b)
				r += d
			}
		}
		return r
	}
	coords := extendUniverse(2)
	p1 := sumMinDistances(coords)
	fmt.Println("Part 1:", p1)
	coords = extendUniverse(1000000)
	p2 := sumMinDistances(coords)
	fmt.Println("Part 2:", p2)
}
