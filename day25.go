package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

type Node struct {
	name       string
	neighbours []*Node
}

func main() {
	nodes := map[string]*Node{}
	getNode := func(name string) *Node {
		node, ok := nodes[name]
		if !ok {
			node = &Node{name, []*Node{}}
			nodes[name] = node
		}
		return node
	}
	edges := [][2]*Node{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		neighbours := strings.Split(parts[1], " ")
		node := getNode(parts[0])
		for _, name := range neighbours {
			node1 := getNode(name)
			node.neighbours = append(node.neighbours, node1)
			node1.neighbours = append(node1.neighbours, node)
			edges = append(edges, [2]*Node{node, node1})
		}
	}
	n := len(nodes)
	nodesArr := make([]*Node, 0, n)
	for _, node := range nodes {
		nodesArr = append(nodesArr, node)
	}
	for {
		freq := map[int]int{}
		for i := 0; i < 1000; i++ {
			startId := rand.Intn(n)
			endId := rand.Intn(n)
			if startId == endId {
				continue
			}
			path := getPath(nodesArr[startId], nodesArr[endId])
			for i := 1; i < len(path); i++ {
				edgeId := getEdgeId(edges, path[i-1], path[i])
				freq[edgeId] += 1
			}
		}
		freqArr := [][2]int{}
		for k, v := range freq {
			freqArr = append(freqArr, [2]int{k, v})
		}
		sort.Slice(freqArr, func(i, j int) bool {
			return freqArr[i][1] > freqArr[j][1]
		})
		exclude := [3]int{freqArr[0][0], freqArr[1][0], freqArr[2][0]}
		group1 := graphSize(nodesArr[0], edges, exclude)
		group2 := n - group1
		if group1*group2 != 0 {
			fmt.Println("Part 1:", group1*group2)
			break
		}
	}
}

func getPath(start, end *Node) []*Node {
	prev := map[*Node]*Node{start: start}
	queue := []*Node{start}
	visited := map[*Node]bool{start: true}
	for len(queue) > 0 {
		new_queue := []*Node{}
		for _, node := range queue {
			for _, neighbour := range node.neighbours {
				if _, ok := visited[neighbour]; ok {
					continue
				}
				visited[neighbour] = true
				prev[neighbour] = node
				new_queue = append(new_queue, neighbour)
			}
		}
		queue = new_queue
	}
	if _, ok := prev[end]; !ok {
		return nil
	}
	path := []*Node{}
	current := end
	for current != start {
		path = append([]*Node{current}, path...)
		current = prev[current]
	}
	return path
}

func getEdgeId(edges [][2]*Node, a, b *Node) int {
	for i, edge := range edges {
		if edge[0] == a && edge[1] == b {
			return i
		}
		if edge[1] == a && edge[0] == b {
			return i
		}
	}
	return -1
}

func graphSize(root *Node, edges [][2]*Node, exclude [3]int) int {
	queue := []*Node{root}
	visited := map[*Node]bool{root: true}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, neighbour := range current.neighbours {
			edgeId := getEdgeId(edges, current, neighbour)
			if edgeId == exclude[0] || edgeId == exclude[1] || edgeId == exclude[2] {
				continue
			}
			if _, ok := visited[neighbour]; ok {
				continue
			}
			visited[neighbour] = true
			queue = append(queue, neighbour)
		}
	}
	return len(visited)
}
