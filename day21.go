package main

import (
	"bufio"
	"os"
	"fmt"
)

type Coord struct {
	x, y int
}

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	n, m := len(lines), len(lines[0])
	var start Coord
	for i, line := range lines {
		for j, ch := range line {
			if ch == 'S' {
				start = Coord{i, j}
			}
		}
	}
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}
	isValid := func(p Coord) bool {
		if p.x < 0 || p.x >= n || p.y < 0 || p.y >= m {
			return false
		}
		if lines[p.x][p.y] == '#' {
			return false
		}
		return true
	}
	move := func(p Coord, validate func(Coord)bool) []Coord {
		r := []Coord{}
		for i := 0; i < 4; i++ {
			next := Coord{p.x + dx[i], p.y + dy[i]}
			if validate(next) {
				r = append(r, next)
			}
		}
		return r
	}
	type tuple struct {
		neu, alt int
	}
	history := []tuple{}
	moveN := func(steps int, validate func(Coord)bool) int {
		visited := map[Coord]bool{}
		queue := []Coord{ start }
		for i := 0; i < steps; i++ {
			newQueueUniq := map[Coord]bool{}
			for _, p := range queue {
				for _, pp := range move(p, validate) {
					newQueueUniq[pp] = true
				}
			}
			neuVisited := 0
			newQueue := []Coord{}
			for p := range newQueueUniq {
				if _, ok := visited[p]; !ok {
					neuVisited++
					visited[p] = true
				}
				newQueue = append(newQueue, p)
			}
			history = append(history, tuple{ neuVisited, len(newQueue) - neuVisited })
			queue = newQueue
		}
		return len(queue)
	}
	p1 := moveN(64, isValid)
	fmt.Println("Part 1:", p1)
	isValidInf := func(p Coord) bool {
		x, y := p.x % n, p.y % m
		if x < 0 {
			x += n
		}
		if y < 0 {
			y += m
		}
		if lines[x][y] == '#' {
			return false
		}
		return true
	}
	newCycleStart := 64
	newCycleLength := 131
	history = []tuple{}
	moveN(500, isValidInf)
	diffs1 := make([]int, newCycleLength)
	diffs2 := make([]int, newCycleLength)
	for i := 0; i < newCycleLength; i++ {
		idx := newCycleStart + i
		diffs1[i] = history[idx].neu - history[idx-1].neu
		diffs2[i] = (history[idx + newCycleLength].neu - history[idx + newCycleLength - 1].neu) - (history[idx].neu - history[idx-1].neu)
	}
	prevPrevResult := history[448].neu + history[448].alt
	prevResult, prevNew := history[449].neu + history[449].alt, history[449].neu
	for i := 450; i < 26501365; i++ {
		idx := (i - newCycleStart) % newCycleLength
		mul := (i - newCycleStart) / newCycleLength
		nextNew := prevNew + (diffs1[idx] + mul * diffs2[idx])
		nextOld := prevPrevResult
		nextResult := nextNew + nextOld
		prevPrevResult, prevResult, prevNew = prevResult, nextResult, nextNew
	}
	fmt.Println("Part 2:", prevResult)
}
