package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arrDiff := func(arr []int) []int {
		r := make([]int, len(arr)-1)
		for i := 1; i < len(arr); i++ {
			r[i-1] = arr[i] - arr[i-1]
		}
		return r
	}
	arrIsZero := func(arr []int) bool {
		for _, value := range arr {
			if value != 0 {
				return false
			}
		}
		return true
	}
	var arrGetNext func([]int) int
	arrGetNext = func(arr []int) int {
		diff := arrDiff(arr)
		if arrIsZero(diff) {
			return arr[len(arr)-1]
		} else {
			return arr[len(arr)-1] + arrGetNext(diff)
		}
	}
	var arrGetPrev func([]int) int
	arrGetPrev = func(arr []int) int {
		diff := arrDiff(arr)
		if arrIsZero(diff) {
			return arr[0]
		} else {
			return arr[0] - arrGetPrev(diff)
		}
	}
	p1 := 0
	p2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		arr := make([]int, len(parts))
		for i, part := range parts {
			value, _ := strconv.Atoi(part)
			arr[i] = value
		}
		p1 += arrGetNext(arr)
		p2 += arrGetPrev(arr)
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
