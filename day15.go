package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ",")
	p1 := 0
	hash := func(s string) int {
		current := 0
		for _, ch := range []byte(s) {
			current += int(ch)
			current *= 17
			current %= 256
		}
		return current
	}
	for _, part := range parts {
		p1 += hash(part)
	}
	fmt.Println("Part 1:", p1)
	type pair struct {
		label string
		num int
	}
	boxes := make([][]pair, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = []pair{}
	}
	for _, part := range parts {
		if part[len(part) - 1] == '-' {
			label := part[:len(part) - 1]
			boxId := hash(label)
			box := boxes[boxId]
			if len(box) > 0 {
				r := []pair{}
				for _, p := range box {
					if p.label != label {
						r = append(r, p)
					}
				}
				boxes[boxId] = r
			}
		} else {
			arr := strings.Split(part, "=")
			label := arr[0]
			num, _ := strconv.Atoi(arr[1])
			boxId := hash(label)
			box := boxes[boxId]
			r := []pair{}
			found := false
			for _, p := range box {
				if p.label == label {
					found = true
					r = append(r, pair{label, num})
				} else {
					r = append(r, p)
				}
			}
			if !found {
				r = append(r, pair{label, num})
			}
			boxes[boxId] = r
		}
	}
	p2 := 0
	for i, box := range boxes {
		for j, lens := range box {
			//fmt.Println((i + 1), (j + 1), lens.num)
			p2 += (i + 1) * (j + 1) * lens.num
		}
	}
	fmt.Println("Part 2:", p2)
}
