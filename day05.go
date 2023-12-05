package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func arrToInt(arr []string) []int {
	r := make([]int, len(arr))
	for i, x := range arr {
		r[i], _ = strconv.Atoi(x)
	}
	return r
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	seedsStr := strings.Split(line, " ")
	seeds := arrToInt(seedsStr[1:])
	type tuple struct {
		x, y, z int
	}
	m := [][]tuple{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scanner.Scan()
			m = append(m, []tuple{})
			continue
		}
		parts := arrToInt(strings.Split(line, " "))
		t := tuple{parts[0], parts[1], parts[2]}
		cur := m[len(m)-1]
		m[len(m)-1] = append(cur, t)
	}
	res := make([]int, len(seeds))
	for i, seed := range seeds {
		for _, arr := range m {
			for _, t := range arr {
				if seed >= t.y && seed <= t.y+t.z {
					diff := seed - t.y
					seed = t.x + diff
					break
				}
			}
		}
		res[i] = seed
	}
	sort.Ints(res)
	p1 := res[0]
	fmt.Println("Part 1:", p1)
	type pair struct {
		x, y int
	}
	hasNoIntersection := func(a pair, t tuple) bool {
		return a.x+a.y <= t.y || a.x >= t.y+t.z
	}
	intersect := func(a pair, t tuple) (in pair, out []pair) {
		if a.x >= t.y && a.x+a.y <= t.y+t.z {
			diff := a.x - t.y
			in.x = t.x + diff
			in.y = a.y
			return
		}
		if a.x < t.y && a.x+a.y <= t.y+t.z {
			out = append(out, pair{a.x, t.y - a.x})
			in.x = t.x
			in.y = a.y - (t.y - a.x)
			return
		}
		if a.x >= t.y && a.x+a.y > t.y+t.z {
			diff := a.x - t.y
			in.x = t.x + diff
			in.y = t.y + t.z - a.x
			out = append(out, pair{t.y + t.z, a.y - in.y})
			return
		}
		in.x = t.x
		in.y = t.z
		out = append(out, pair{a.x, t.y - a.x})
		out = append(out, pair{t.y + t.z, a.y - (t.y - a.x) - t.z})
		return
	}
	queue := make([]pair, len(seeds)/2)
	for i := 0; i < len(seeds)/2; i++ {
		queue[i] = pair{seeds[2*i+0], seeds[2*i+1]}
	}
	for _, arr := range m {
		match := []pair{}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			found := false
			for _, t := range arr {
				if hasNoIntersection(p, t) {
					continue
				}
				in, out := intersect(p, t)
				match = append(match, in)
				for _, x := range out {
					queue = append(queue, x)
				}
				found = true
				break
			}
			if !found {
				match = append(match, p)
			}
		}
		queue = match
	}
	sort.Slice(queue, func(i, j int) bool {
		return queue[i].x < queue[j].x
	})
	p2 := queue[0].x
	fmt.Println("Part 2:", p2)
}
