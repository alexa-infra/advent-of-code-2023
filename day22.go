package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Coord struct {
	x, y, z int
}

type Cube struct {
	a, b Coord
}

func intersect(c1, c2 *Cube, dz int) bool {
	return ((c1.b.z+dz) >= c2.a.z &&
		(c1.a.z+dz) <= c2.b.z &&
		c1.b.x >= c2.a.x &&
		c1.a.x <= c2.b.x &&
		c1.b.y >= c2.a.y &&
		c1.a.y <= c2.b.y)
}

func main() {
	cubes := []Cube{}
	re := regexp.MustCompile(`(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		sx, sy, sz := match[1], match[2], match[3]
		x, _ := strconv.Atoi(sx)
		y, _ := strconv.Atoi(sy)
		z, _ := strconv.Atoi(sz)
		a := Coord{x, y, z}
		sx, sy, sz = match[4], match[5], match[6]
		x, _ = strconv.Atoi(sx)
		y, _ = strconv.Atoi(sy)
		z, _ = strconv.Atoi(sz)
		b := Coord{x, y, z}
		cubes = append(cubes, Cube{a, b})
	}
	sort.Slice(cubes, func(i, j int) bool {
		return cubes[i].a.z < cubes[j].a.z
	})
	fall := func(cubes []Cube) int {
		counter := 0
		for i := range cubes {
			a := &cubes[i]
			hasFell := false
			for a.a.z > 1 {
				canFall := true
				for j := i - 1; j >= 0; j-- {
					b := &cubes[j]
					if intersect(a, b, -1) {
						canFall = false
						break
					}
				}
				if canFall {
					a.a.z--
					a.b.z--
					hasFell = true
				} else {
					break
				}
			}
			if hasFell {
				counter++
			}
		}
		return counter
	}
	fall(cubes)
	check := func(cubes []Cube) (int, int) {
		tmp := make([]Cube, len(cubes))
		counter, allMoves := 0, 0
		for i := range cubes {
			copy(tmp, cubes)
			tmp[i] = Cube{Coord{0, 0, 0}, Coord{0, 0, 0}}
			moved := fall(tmp)
			if moved == 0 {
				counter += 1
			} else {
				allMoves += moved
			}
		}
		return counter, allMoves
	}
	p1, p2 := check(cubes)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
