package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	pulseLow    = "LOW"
	pulseHigh   = "HIGH"
	flipFlopOn  = "ON"
	flipFlopOff = "OFF"
)

type Node struct {
	name     string
	kind     byte
	dest     []*Node
	state    string
	stateMap map[string]string
}

func MakeNode(name string) *Node {
	return &Node{name, byte(0), []*Node{}, flipFlopOff, map[string]string{}}
}

type Pulse struct {
	kind string
	src  *Node
	dest *Node
}

func (node *Node) receive(p Pulse, emit func(Pulse)) {
	emitForChildren := func(pulseType string) {
		for _, dest := range node.dest {
			emit(Pulse{pulseType, node, dest})
		}
	}
	if node.kind == '%' {
		if p.kind == pulseLow {
			if node.state == flipFlopOn {
				emitForChildren(pulseLow)
				node.state = flipFlopOff
			} else if node.state == flipFlopOff {
				emitForChildren(pulseHigh)
				node.state = flipFlopOn
			} else {
				fmt.Println("Err")
			}
		}
		return
	}
	if node.kind == '&' {
		node.stateMap[p.src.name] = p.kind
		for _, v := range node.stateMap {
			if v == pulseLow {
				emitForChildren(pulseHigh)
				return
			}
		}
		emitForChildren(pulseLow)
		return
	}
	if node.name == "broadcaster" {
		emitForChildren(p.kind)
		return
	}
}

func main() {
	re := regexp.MustCompile(`(.*) -> (.*)`)
	scanner := bufio.NewScanner(os.Stdin)
	nodes := map[string]*Node{}
	getOrCreateNode := func(name string) *Node {
		node, ok := nodes[name]
		if !ok {
			node = MakeNode(name)
			nodes[name] = node
		}
		return node
	}
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		name := match[1]
		kind := byte(0)
		if name[0] == '%' || name[0] == '&' {
			kind = name[0]
			name = name[1:]
		}
		node := getOrCreateNode(name)
		if kind != 0 {
			node.kind = kind
		}
		destNames := strings.Split(match[2], ", ")
		for _, dname := range destNames {
			node1 := getOrCreateNode(dname)
			node.dest = append(node.dest, node1)
			node1.stateMap[name] = pulseLow
		}
	}
	queue := []Pulse{}
	lowCounter, highCounter := int(0), int(0)
	emit := func(p Pulse) {
		if p.kind == pulseLow {
			lowCounter++
		}
		if p.kind == pulseHigh {
			highCounter++
		}
		queue = append(queue, p)
	}
	processQueue := func() {
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			cur.dest.receive(cur, emit)
		}
	}
	broadcaster := nodes["broadcaster"]
	for i := 0; i < 1000; i++ {
		emit(Pulse{pulseLow, nil, broadcaster})
		processQueue()
	}
	p1 := lowCounter * highCounter
	fmt.Println("Part 1:", p1)
	for _, node := range nodes {
		node.state = flipFlopOff
		for src := range node.stateMap {
			node.stateMap[src] = pulseLow
		}
	}
	emit2 := func(p Pulse) {
		queue = append(queue, p)
	}
	diff := map[string]map[string]int{}
	for _, node := range nodes {
		if node.kind == '&' && len(node.dest) > 2 {
			diff[node.name] = copyMap(node.stateMap)
		}
	}
	for i := 0; i < 10000; i++ {
		for kk, mm := range diff {
			for k, v := range mm {
				if v == 0 && nodes[kk].stateMap[k] == pulseHigh {
					mm[k] = i
				}
			}
		}
		emit2(Pulse{pulseLow, nil, broadcaster})
		processQueue()
	}
	arr := []int{}
	for _, mm := range diff {
		num := 0
		for _, v := range mm {
			num |= v
		}
		arr = append(arr, num)
	}
	p2 := LCMArray(arr)
	fmt.Println("Part 2:", p2)
}

func copyMap(m map[string]string) map[string]int {
	c := map[string]int{}
	for k := range m {
		c[k] = 0
	}
	return c
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
