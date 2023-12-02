package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile(`Game (\d+): (.*)`)
	re2 := regexp.MustCompile(`(\d+) (blue|red|green)([,;])? ?`)
	p1 := 0
	p2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		gameId, _ := strconv.Atoi(match[1])
		match2 := re2.FindAllStringSubmatch(match[2], -1)
		blue, red, green := 0, 0, 0
		for _, r := range match2 {
			val, _ := strconv.Atoi(r[1])
			if r[2] == "blue" {
				blue = max(blue, val)
			}
			if r[2] == "red" {
				red = max(red, val)
			}
			if r[2] == "green" {
				green = max(green, val)
			}
		}
		if red <= 12 && green <= 13 && blue <= 14 {
			p1 += gameId
		}
		p2 += blue * red * green
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
