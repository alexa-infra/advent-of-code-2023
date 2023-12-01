package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	sum := 0
	re := regexp.MustCompile(`[0-9]`)
	part1 := func(line string) int {
		match := re.FindAllString(line, -1)
		if len(match) == 0 {
			return 0
		}
		first, last := match[0], match[len(match)-1]
		firstInt, _ := strconv.Atoi(first)
		lastInt, _ := strconv.Atoi(last)
		return firstInt*10 + lastInt
	}
	sum2 := 0
	re2 := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[0-9])`)
	part2 := func(line string) int {
		line = strings.ReplaceAll(line, "oneight", "one.eight")
		line = strings.ReplaceAll(line, "twone", "two.one")
		line = strings.ReplaceAll(line, "threeight", "three.eight")
		line = strings.ReplaceAll(line, "fiveight", "five.eight")
		line = strings.ReplaceAll(line, "sevenine", "seven.nine")
		line = strings.ReplaceAll(line, "eightwo", "eight.two")
		line = strings.ReplaceAll(line, "eighthree", "eight.three")
		line = strings.ReplaceAll(line, "nineight", "nine.eight")

		match := re2.FindAllString(line, -1)
		if len(match) == 0 {
			return 0
		}
		first, last := match[0], match[len(match)-1]
		firstInt, ok := m[first]
		if !ok {
			firstInt, _ = strconv.Atoi(first)
		}
		lastInt, ok := m[last]
		if !ok {
			lastInt, _ = strconv.Atoi(last)
		}
		return firstInt*10 + lastInt
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		sum += part1(line)
		sum2 += part2(line)
	}
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}
