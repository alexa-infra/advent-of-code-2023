package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	m := map[rune]int{
		'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8, '9': 7,
		'8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0,
	}
	type tuple struct {
		orig string
		hand []int
		kind int
		bid  int
	}
	freqToKind := func(freq map[rune]int) int {
		arr := []int{}
		for _, v := range freq {
			arr = append(arr, v)
		}
		kind := 0
		switch len(freq) {
		case 5: // high card
			kind = 0
		case 4: // one pair
			kind = 1
		case 3: // two pairs or three of a kind
			x1, x2, x3 := arr[0], arr[1], arr[2]
			if x1 == 3 || x2 == 3 || x3 == 3 {
				kind = 3
			} else {
				kind = 2
			}
		case 2: // full house or four of a kind
			x1, x2 := arr[0], arr[1]
			if x1 == 4 || x2 == 4 {
				kind = 5
			} else {
				kind = 4
			}
		case 1: // five of a kind
			kind = 6
		default:
			fmt.Println("Err")
		}
		return kind
	}
	cards := []tuple{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		hand := make([]int, 5)
		freq := map[rune]int{}
		for i, ch := range parts[0] {
			hand[i] = m[ch]
			freq[ch] += 1
		}
		kind := freqToKind(freq)
		cards = append(cards, tuple{parts[0], hand, kind, bid})
	}
	sortFunc := func(i, j int) bool {
		a, b := cards[i], cards[j]
		if a.kind != b.kind {
			return a.kind < b.kind
		}
		for i := 0; i < 5; i++ {
			if a.hand[i] != b.hand[i] {
				return a.hand[i] < b.hand[i]
			}
		}
		return true
	}
	sort.Slice(cards, sortFunc)
	p1 := 0
	for i, card := range cards {
		p1 += (i + 1) * card.bid
	}
	fmt.Println("Part 1:", p1)
	m['J'] = -1
	for j, card := range cards {
		hand := make([]int, 5)
		freq := map[rune]int{}
		for i, ch := range card.orig {
			hand[i] = m[ch]
			freq[ch] += 1
		}
		v, ok := freq['J']
		if ok {
			delete(freq, 'J')
			arr := []int{}
			for _, v := range freq {
				arr = append(arr, v)
			}
			sort.Ints(arr)
			if len(arr) == 0 {
				freq['J'] = v
			} else {
				max := arr[len(arr)-1]
				for k, vv := range freq {
					if vv == max {
						freq[k] += v
						break
					}
				}
			}
		}
		kind := freqToKind(freq)
		cards[j].hand = hand
		cards[j].kind = kind
	}
	sort.Slice(cards, sortFunc)
	p2 := 0
	for i, card := range cards {
		p2 += (i + 1) * card.bid
	}
	fmt.Println("Part 2:", p2)
}
