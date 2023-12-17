package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	n, m := len(lines), len(lines[0])
	type state struct {
		x, y, cost int
		dir, dirN  byte
	}
	valid := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < n && y < m
	}
	toEnd := func(x, y int) int {
		return abs(x-n+1) + abs(y-m+1)
	}
	run := func(getNext func(state) []state) int {
		cache := map[state]int{}
		queue := []state{state{0, 0, 0, 0, 0}}
		hasChanged := false
		for len(queue) > 0 {
			if hasChanged {
				sort.Slice(queue, func(i, j int) bool {
					d1 := queue[i].cost + toEnd(queue[i].x, queue[i].y)
					d2 := queue[j].cost + toEnd(queue[j].x, queue[j].y)
					if d1 == d2 {
						return queue[i].cost < queue[j].cost
					}
					return d1 < d2
				})
				hasChanged = false
			}
			//fmt.Println(len(queue), queue[0].cost)
			cur := queue[0]
			if cur.x == n-1 && cur.y == m-1 {
				return cur.cost
			}
			queue = queue[1:]
			nextStates := getNext(cur)
			for _, next := range nextStates {
				cc := state{next.x, next.y, 0, next.dir, next.dirN}
				v, ok := cache[cc]
				if !ok || next.cost < v {
					cache[cc] = next.cost
					queue = append(queue, next)
					hasChanged = true
				}
			}
		}
		return -1
	}
	directions := []byte{'N', 'S', 'E', 'W'}
	dx := map[byte]int{'N': -1, 'S': 1, 'E': 0, 'W': 0}
	dy := map[byte]int{'N': 0, 'S': 0, 'E': 1, 'W': -1}
	reverse := map[byte]byte{'N': 'S', 'S': 'N', 'E': 'W', 'W': 'E'}
	isReverse := func(a, b byte) bool {
		reverseDir, ok := reverse[a]
		return ok && reverseDir == b
	}
	p1 := run(func(cur state) []state {
		r := []state{}
		mustTurn := cur.dirN >= 3
		for _, dir := range directions {
			if isReverse(cur.dir, dir) {
				continue
			}
			if cur.dir == dir && mustTurn {
				continue
			}
			x := cur.x + dx[dir]
			y := cur.y + dy[dir]
			if !valid(x, y) {
				continue
			}
			cost := cur.cost + int(lines[x][y]-'0')
			dirN := byte(1)
			if cur.dir == dir {
				dirN = cur.dirN + 1
			}
			next := state{x, y, cost, dir, dirN}
			r = append(r, next)
		}
		return r
	})
	fmt.Println("Part 1:", p1)
	p2 := run(func(cur state) []state {
		r := []state{}
		canTurn := cur.dirN >= 4
		mustTurn := cur.dirN == 10
		for _, dir := range directions {
			if isReverse(cur.dir, dir) {
				continue
			}
			if cur.dir == dir && mustTurn {
				continue
			}
			if cur.dir != 0 && cur.dir != dir && !canTurn {
				continue
			}
			x := cur.x + dx[dir]
			y := cur.y + dy[dir]
			if !valid(x, y) {
				continue
			}
			cost := cur.cost + int(lines[x][y]-'0')
			dirN := byte(1)
			if cur.dir == dir {
				dirN = cur.dirN + 1
			}
			next := state{x, y, cost, dir, dirN}
			r = append(r, next)
		}
		return r
	})
	fmt.Println("Part 2:", p2)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
