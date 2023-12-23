package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}

type Path map[Coord]int

func copyPath(p Path) Path {
	r := Path{}
	for k, v := range p {
		r[k] = v
	}
	return r
}

func getLastPos(p Path) Coord {
	n := len(p)
	for k, v := range p {
		if v == n {
			return k
		}
	}
	return Coord{}
}

type Edge struct {
	dest Coord
	dist int
}

type Graph map[Coord][]Edge

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	n, m := len(lines), len(lines[0])
	var start, end Coord
	for i := 0; i < m; i++ {
		if lines[0][i] == '.' {
			start = Coord{0, i}
		}
		if lines[n-1][i] == '.' {
			end = Coord{n - 1, i}
		}
	}
	dirs := [][]int{[]int{-1, 0}, []int{0, 1}, []int{1, 0}, []int{0, -1}}
	validPosition := func(pp Coord, dir []int, slopes bool) bool {
		if pp.x < 0 || pp.x >= n || pp.y < 0 || pp.y >= m {
			return false
		}
		ch := lines[pp.x][pp.y]
		if ch == '#' {
			return false
		}
		if slopes {
			if ch == '>' && dir[1] != 1 {
				return false
			}
			if ch == '<' && dir[1] != -1 {
				return false
			}
			if ch == '^' && dir[0] != -1 {
				return false
			}
			if ch == 'v' && dir[0] != 1 {
				return false
			}
		}
		return true
	}
	waypoints := map[Coord][]Coord{}
	for i, line := range lines {
		for j := range line {
			p := Coord{i, j}
			if !validPosition(p, nil, false) {
				continue
			}
			ways := []Coord{}
			for _, d := range dirs {
				pp := Coord{i + d[0], j + d[1]}
				if !validPosition(pp, nil, false) {
					continue
				}
				ways = append(ways, pp)
			}
			if len(ways) > 2 {
				waypoints[p] = ways
			}
		}
	}
	waypoints[start] = []Coord{Coord{start.x + 1, start.y}}
	waypoints[end] = []Coord{Coord{end.x - 1, end.y}}

	makePath := func(pos1, pos2 Coord, slopes bool) (Coord, int) {
		path := Path{pos1: 1, pos2: 2}
		current := pos2
		for {
			if _, ok := waypoints[current]; ok {
				return current, len(path) - 1
			}
			found := false
			for _, d := range dirs {
				if !validPosition(current, d, slopes) {
					continue
				}
				pp := Coord{current.x + d[0], current.y + d[1]}
				if !validPosition(pp, d, slopes) {
					continue
				}
				if _, ok := path[pp]; ok {
					continue
				}
				path[pp] = len(path) + 1
				current = pp
				found = true
				break
			}
			if !found {
				break
			}
		}
		return Coord{}, 0
	}
	buildGraph := func(slopes bool) Graph {
		edges := Graph{}
		for p, nextArr := range waypoints {
			arr := []Edge{}
			for _, next := range nextArr {
				pp, plen := makePath(p, next, slopes)
				if plen > 0 {
					arr = append(arr, Edge{pp, plen})
				}
			}
			edges[p] = arr
		}
		return edges
	}
	var longest func(Graph, Path) (bool, int)
	longest = func(graph Graph, path Path) (bool, int) {
		current := getLastPos(path)
		if current == end {
			return true, 0
		}
		maxDist := 0
		nextArr := graph[current]
		found := false
		for _, next := range nextArr {
			nextPos := next.dest
			if _, ok := path[nextPos]; ok {
				continue
			}
			newPath := copyPath(path)
			newPath[nextPos] = len(newPath) + 1
			ff, newDist := longest(graph, newPath)
			if ff && newDist + next.dist > maxDist {
				found = true
				maxDist = newDist + next.dist
			}
		}
		return found, maxDist
	}
	solve := func(slopes bool) int {
		graph := buildGraph(slopes)
		_, r := longest(graph, Path{start: 1})
		return r
	}
	p1 := solve(true)
	fmt.Println("Part 1:", p1)
	p2 := solve(false)
	fmt.Println("Part 2:", p2)
}
