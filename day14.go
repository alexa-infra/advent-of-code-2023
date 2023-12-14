package main

import (
	"bufio"
	"os"
	"fmt"
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
	data := map[pair]byte{}
	stones := []pair{}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			p := pair{ n - i, j + 1 }
			ch := lines[i][j]
			data[p] = ch
			if ch == 'O' {
				stones = append(stones, p)
			}
		}
	}
	force := func(dx, dy int) {
		for {
			moved := false
			for i, p := range stones {
				next := pair{p.x + dx, p.y + dy}
				ch2, ok := data[next]
				if ok && ch2 == '.' {
					data[next] = 'O'
					data[p] = '.'
					stones[i] = next
					moved = true
				}
			}
			if !moved {
				break
			}
		}
	}
	calc := func() int {
		r := 0
		for _, p := range stones {
			r += p.x
		}
		return r
	}
	force(1, 0)
	p1 := calc()
	fmt.Println("Part 1:", p1)
	force(0, -1)
	force(-1, 0)
	force(0, 1)
	uniq := map[int]int{}
	uniq[calc()] += 1
	start := false
	startId := 0
	cycle := []int{}
	matches := []bool{}
	for i := 2; i < 1000; i++ {
		force(1, 0)
		force(0, -1)
		force(-1, 0)
		force(0, 1)
		value := calc()
		uniq[value] += 1
		if !start && uniq[value] >= 10 {
			start = true
			startId = i
		}
		if start {
			countMatches := 0
			for j, match := range matches {
				if match {
					countMatches++
				} else if cycle[j] == value {
					matches[j] = true
					countMatches++
					break
				} else {
					for k := range matches {
						matches[k] = false
					}
					countMatches = 0
					break
				}
			}
			if countMatches >= 10 {
				cycle = cycle[:len(cycle) - countMatches + 1]
				break
			}
			cycle = append(cycle, value)
			matches = append(matches, false)
		}
	}
	idx := (1000000000 - startId) % len(cycle)
	fmt.Println("Part 2:", cycle[idx])
}
