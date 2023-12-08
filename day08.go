package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	re := regexp.MustCompile(`(...) = \((...), (...)\)`)
	scanner.Scan()
	directions := scanner.Text()
	type Node struct {
		name        string
		left, right *Node
	}
	nodes := map[string]*Node{}
	makeNode := func(name string) *Node {
		node, ok := nodes[name]
		if !ok {
			node = &Node{name, nil, nil}
			nodes[name] = node
		}
		return node
	}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		match := re.FindStringSubmatch(line)
		node := makeNode(match[1])
		node.left = makeNode(match[2])
		node.right = makeNode(match[3])
	}
	start := nodes["AAA"]
	end := nodes["ZZZ"]
	current := start
	steps := 0
	pos := 0
	for current != end {
		ch := directions[pos]
		pos = (pos + 1) % len(directions)
		if ch == 'L' {
			current = current.left
		} else if ch == 'R' {
			current = current.right
		}
		steps++
	}
	fmt.Println("Part 1:", steps)
	ghosts := []*Node{}
	for _, v := range nodes {
		if v.name[2] == 'A' {
			ghosts = append(ghosts, v)
		}
	}
	steps = 0
	pos = 0
	last := map[int]int{}
	freq := map[int]int{}
	for {
		cc := 0
		for i, ghost := range ghosts {
			if ghost.name[2] == 'Z' {
				cc++
				pp, ok := last[i]
				if ok {
					freq[i] = steps - pp
				}
				last[i] = steps
			}
		}
		if len(freq) == len(ghosts) {
			//fmt.Println(cc, "/", len(ghosts), "at", steps, freq, last)
			break
		}
		ch := directions[pos]
		pos = (pos + 1) % len(directions)
		for i, ghost := range ghosts {
			if ch == 'L' {
				ghosts[i] = ghost.left
			} else if ch == 'R' {
				ghosts[i] = ghost.right
			}
		}
		steps++
	}
	freqArr := make([]int, len(freq))
	for i, v := range freq {
		freqArr[i] = v
	}
	p2 := LCMArray(freqArr)
	fmt.Println("Part 2:", p2)
}

// Greatest Common Divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func LCMArray(arr []int) int {
	a, b := arr[0], arr[1]
	result := LCM(a, b)

	for i := 2; i < len(arr); i++ {
		result = LCM(result, arr[i])
	}

	return result
}
