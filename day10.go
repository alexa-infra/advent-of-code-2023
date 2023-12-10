package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	type pair struct {
		x, y int
	}
	tubes := map[pair]rune{}
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	n, m := len(lines), len(lines[0])
	var start pair
	for i, line := range lines {
		for j, ch := range line {
			p := pair{j, i}
			if ch != '.' {
				tubes[p] = ch
			}
			if ch == 'S' {
				start = p
			}
		}
	}

	directions := map[rune]map[rune]rune{
		'E': map[rune]rune{
			'-': 'E',
			'J': 'N',
			'7': 'S',
		},
		'W': map[rune]rune{
			'-': 'W',
			'L': 'N',
			'F': 'S',
		},
		'S': map[rune]rune{
			'|': 'S',
			'L': 'E',
			'J': 'W',
		},
		'N': map[rune]rune{
			'|': 'N',
			'7': 'W',
			'F': 'E',
		},
	}
	directions2 := map[rune][]rune{
		'|': []rune{'N', 'S'},
		'-': []rune{'W', 'E'},
		'L': []rune{'N', 'E'},
		'J': []rune{'N', 'W'},
		'7': []rune{'S', 'W'},
		'F': []rune{'S', 'E'},
	}
	canMove := func(p pair, dir rune) (pair, rune, bool) {
		var next pair
		switch dir {
		case 'N':
			next = pair{p.x, p.y - 1}
		case 'S':
			next = pair{p.x, p.y + 1}
		case 'W':
			next = pair{p.x - 1, p.y}
		case 'E':
			next = pair{p.x + 1, p.y}
		}
		ch, ok := tubes[next]
		if !ok {
			return p, dir, false
		}
		m := directions[dir]
		nextDir, ok := m[ch]
		if !ok {
			return p, dir, false
		}
		return next, nextDir, true
	}
	p1 := 0
	for ch, dd := range directions2 {
		tubes[start] = ch

		pos1, dir1, ok1 := canMove(start, dd[0])
		pos2, dir2, ok2 := canMove(start, dd[1])
		steps := 1
		for ok1 && ok2 && pos1 != pos2 {
			pos1, dir1, ok1 = canMove(pos1, dir1)
			pos2, dir2, ok2 = canMove(pos2, dir2)
			steps++
		}
		if ok1 && ok2 && pos1 == pos2 {
			p1 = steps
			break
		}
	}
	fmt.Println("Part 1:", p1)
	path := map[pair]bool{}
	pos, dir := start, directions2[tubes[start]][0]
	for {
		if _, ok := path[pos]; ok {
			break
		}
		path[pos] = true
		pos, dir, _ = canMove(pos, dir)
	}
	visited := map[pair]bool{}
	pos, dir = start, directions2[tubes[start]][0]
	inOut := map[pair]bool{}
	setInOut := func(x, y int, value bool) {
		p := pair{x, y}
		if x < 0 || y < 0 || x >= m || y >= n {
			return
		}
		if _, ok := path[p]; ok {
			return
		}
		val, ok := inOut[p]
		if ok && val != value {
			fmt.Println("Err")
		} else {
			inOut[p] = value
		}
	}
	for {
		if _, ok := visited[pos]; ok {
			break
		}
		visited[pos] = true
		ch := tubes[pos]
		if ch == '|' {
			right := dir == 'N'
			setInOut(pos.x+1, pos.y, !right)
			setInOut(pos.x-1, pos.y, right)
		} else if ch == '-' {
			right := dir == 'E'
			setInOut(pos.x, pos.y+1, !right)
			setInOut(pos.x, pos.y-1, right)
		} else if ch == 'J' {
			right := dir == 'N'
			setInOut(pos.x+1, pos.y, !right)
			setInOut(pos.x, pos.y+1, !right)
		} else if ch == 'L' {
			right := dir == 'N'
			setInOut(pos.x-1, pos.y, right)
			setInOut(pos.x, pos.y+1, right)
		} else if ch == '7' {
			right := dir == 'S'
			setInOut(pos.x, pos.y-1, right)
			setInOut(pos.x+1, pos.y, right)
		} else if ch == 'F' {
			right := dir == 'E'
			setInOut(pos.x-1, pos.y, right)
			setInOut(pos.x, pos.y-1, right)
		}
		pos, dir, _ = canMove(pos, dir)
	}
	p2 := 0
	queue := []pair{}
	for p, val := range inOut {
		if val {
			queue = append(queue, p)
			p2++
		}
	}
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			pp := pair{p.x + dx[i], p.y + dy[i]}
			if pp.x < 0 || pp.y < 0 || pp.x >= m || pp.y >= n {
				continue
			}
			_, ok := inOut[pp]
			if ok {
				continue
			}
			_, ok = path[pp]
			if ok {
				continue
			}
			inOut[pp] = true
			queue = append(queue, pp)
			p2++
		}
	}
	fmt.Println("Part 2:", p2)
	//for i, line := range lines {
	//	for j, _ := range line {
	//		p := pair{j, i}
	//		if _, ok := path[p]; ok {
	//			fmt.Print(" ")
	//		} else {
	//			if v, ok := inOut[p]; ok && v {
	//				fmt.Print("O")
	//			} else {
	//				fmt.Print(".")
	//			}
	//		}
	//	}
	//	fmt.Println()
	//}
}
