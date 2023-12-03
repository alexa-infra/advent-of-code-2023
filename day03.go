package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	m := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, line)
	}
	type coord struct {
		x, y int
	}
	nums := []coord{}
	symb := map[coord]rune{}
	for i, line := range m {
		for j, ch := range line {
			c := coord{i, j}
			if ch >= '0' && ch <= '9' {
				nums = append(nums, c)
				continue
			}
			if ch == '.' {
				continue
			}
			symb[c] = ch
		}
	}
	mapped := map[coord][]int{}
	p1 := 0
	cur := []byte{}
	adj := map[coord]bool{}
	for p, c := range nums {
		cur = append(cur, m[c.x][c.y])
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				cc := coord{c.x + i, c.y + j}
				_, ok := symb[cc]
				if ok {
					adj[cc] = true
				}
			}
		}
		reset := false
		if p+1 < len(nums) {
			cc := nums[p+1]
			if cc.x != c.x || cc.y != c.y+1 {
				reset = true
			}
		} else {
			reset = true
		}
		if reset {
			curStr := string(cur)
			curNum, _ := strconv.Atoi(curStr)
			if len(adj) > 0 {
				p1 += curNum
				for x := range adj {
					arr, ok := mapped[x]
					if !ok {
						arr = []int{}
					}
					mapped[x] = append(arr, curNum)
				}
			}
			adj = map[coord]bool{}
			cur = []byte{}
		}
	}
	fmt.Println("Part 1:", p1)
	p2 := 0
	for c, arr := range mapped {
		s := symb[c]
		if s == '*' && len(arr) == 2 {
			ratio := arr[0] * arr[1]
			p2 += ratio
		}
	}
	fmt.Println("Part 2:", p2)
}
