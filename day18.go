package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Move struct {
	dir string
	num int
}

type Coord struct {
	x, y int
}

type Line struct {
	start, end Coord
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func shoelace(lines []Line) int {
	s := 0
	for _, line := range lines {
		s += (line.start.y + line.end.y) * (line.start.x - line.end.x)
		if line.start.x == line.end.x {
			s += abs(line.start.y - line.end.y)
		} else {
			s += abs(line.start.x - line.end.x)
		}
	}
	return s/2 + 1
}

func solve(moves []Move) int {
	lines := []Line{}
	current := Coord{0, 0}
	dx := map[string]int{"U": 1, "D": -1, "R": 0, "L": 0}
	dy := map[string]int{"U": 0, "D": 0, "R": 1, "L": -1}
	for _, move := range moves {
		start := current
		x := dx[move.dir]
		y := dy[move.dir]
		current.x += x * move.num
		current.y += y * move.num
		end := current
		lines = append(lines, Line{start, end})
	}
	return shoelace(lines)
}

func main() {
	re := regexp.MustCompile(`([UDRL]) (\d+) \(#(.*)\)`)
	scanner := bufio.NewScanner(os.Stdin)
	text := []string{}
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	moves := []Move{}
	for _, line := range text {
		match := re.FindStringSubmatch(line)
		dir := match[1]
		num, _ := strconv.Atoi(match[2])
		moves = append(moves, Move{dir, num})
	}
	p1 := solve(moves)
	fmt.Println("Part 1:", p1)

	dirMap := map[byte]string{ '0': "R", '1': "D", '2': "L", '3': "U" }
	moves = []Move{}
	for _, line := range text {
		match := re.FindStringSubmatch(line)
		color := match[3]
		dir := dirMap[color[5]]
		num, _ := strconv.ParseInt(color[:5], 16, 32)
		moves = append(moves, Move{dir, int(num)})
	}
	p2 := solve(moves)
	fmt.Println("Part 2:", p2)
}
