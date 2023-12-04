package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile(`Card +(\d+): (.*) \| (.*)`)
	re2 := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(os.Stdin)
	p1 := 0
	own := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		idx, _ := strconv.Atoi(match[1])
		own[idx] += 1
		m1 := re2.FindAllString(match[2], -1)
		m2 := re2.FindAllString(match[3], -1)
		winning := map[string]bool{}
		for _, s := range m1 {
			winning[s] = true
		}
		points := 0
		cc := 0
		for _, s := range m2 {
			if _, ok := winning[s]; ok {
				cc++
				if points == 0 {
					points = 1
				} else {
					points = points << 1
				}
			}
		}
		p1 += points
		for i := 0; i < cc; i++ {
			own[idx+i+1] += own[idx]
		}
	}
	fmt.Println("Part 1:", p1)
	p2 := 0
	for _, v := range own {
		p2 += v
	}
	fmt.Println("Part 2:", p2)
}
