package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	cache := map[string]int{}
	var crack func ([]byte, []int) int
	crack = func (line []byte, hash []int) int {
		key := fmt.Sprintf("%v %v", line, hash)
		value, ok := cache[key]
		if ok {
			return value
		}
		if len(hash) == 0 {
			for _, ch := range line {
				if ch == '#' {
					cache[key] = 0
					return 0
				}
			}
			cache[key] = 1
			return 1
		}
		first := hash[0]
		rest := hash[1:]
		after := sum(rest) + len(rest)
		count := 0
		for i := 0; i <= len(line) - after - first; i++ {
			test := make([]byte, 0, i + first + 1)
			for j := 0; j < i; j++ {
				test = append(test, '.')
			}
			for j := 0; j < first; j++ {
				test = append(test, '#')
			}
			test = append(test, '.')
			match := true
			for j := 0; j < min(len(line), len(test)); j++ {
				if test[j] != line[j] && line[j] != '?' {
					match = false
					break
				}
			}
			if match {
				if len(test) < len(line) {
					count += crack(line[len(test):], rest)
				} else {
					count += crack([]byte{}, rest)
				}
			}
		}
		cache[key] = count
		return count
	}
	p1 := 0
	p2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		numPart := strings.Split(parts[1], ",")
		nums := make([]int, len(numPart))
		for i, p := range numPart {
			v, _ := strconv.Atoi(p)
			nums[i] = v
		}
		pattern := parts[0]
		p1 += crack([]byte(pattern), nums)
		extPattern := strings.Join([]string{ pattern, pattern, pattern, pattern, pattern }, "?")
		extNums := make([]int, len(nums) * 5)
		for i := 0; i < 5; i++ {
			for j, n := range nums {
				extNums[len(nums) * i + j] = n
			}
		}
		p2 += crack([]byte(extPattern), extNums)
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func sum(arr []int) int {
	s := 0
	for _, x := range arr {
		s += x
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
