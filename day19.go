package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pair struct {
	left, right int
}

type Op struct {
	next string
	desc byte
	sign byte
	num  int
}

func main() {
	re1 := regexp.MustCompile(`([a-z]+)\{(.*)\}`)
	re3 := regexp.MustCompile(`([xmas])([\<\>])(.*)\:(.*)`)
	scanner := bufio.NewScanner(os.Stdin)
	flows := map[string][]Op{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		match := re1.FindStringSubmatch(line)
		name := match[1]
		text := match[2]
		parts := strings.Split(text, ",")
		ops := []Op{}
		for _, part := range parts {
			if re3.MatchString(part) {
				m := re3.FindStringSubmatch(part)
				ival, _ := strconv.Atoi(m[3])
				ops = append(ops, Op{m[4], m[1][0], m[2][0], ival})
			} else {
				ops = append(ops, Op{part, 0, 0, 0})
			}
		}
		flows[name] = ops
	}
	data := []map[byte]int{}
	re2 := regexp.MustCompile(`\{x=(.*),m=(.*),a=(.*),s=(.*)\}`)
	for scanner.Scan() {
		line := scanner.Text()
		match := re2.FindStringSubmatch(line)
		x, _ := strconv.Atoi(match[1])
		m, _ := strconv.Atoi(match[2])
		a, _ := strconv.Atoi(match[3])
		s, _ := strconv.Atoi(match[4])
		data = append(data, map[byte]int{'x': x, 'm': m, 'a': a, 's': s})
	}
	var follow1 func(string, map[byte]int) int
	follow1 = func(name string, values map[byte]int) int {
		if name == "A" {
			return values['x'] + values['m'] + values['a'] + values['s']
		}
		if name == "R" {
			return 0
		}
		r := 0
		ops := flows[name]
		for _, op := range ops {
			if op.sign == 0 {
				r += follow1(op.next, values)
			} else if op.sign == '<' {
				if values[op.desc] < op.num {
					r += follow1(op.next, values)
					break
				}
			} else if op.sign == '>' {
				if values[op.desc] > op.num {
					r += follow1(op.next, values)
					break
				}
			}
		}
		return r
	}
	p1 := 0
	for _, v := range data {
		p1 += follow1("in", v)
	}
	fmt.Println("Part 1:", p1)
	var follow func(string, map[byte]pair) int
	follow = func(name string, values map[byte]pair) int {
		if name == "R" {
			return 0
		}
		if name == "A" {
			p := 1
			for _, v := range values {
				p *= (v.right - v.left + 1)
			}
			return p
		}
		r := 0
		ops := flows[name]
		for _, op := range ops {
			if op.sign == 0 {
				r += follow(op.next, values)
			} else if op.sign == '<' {
				v := values[op.desc]
				cc := clone(values)
				cc[op.desc] = pair{v.left, op.num - 1}
				r += follow(op.next, cc)
				values[op.desc] = pair{op.num, v.right}
			} else if op.sign == '>' {
				v := values[op.desc]
				cc := clone(values)
				cc[op.desc] = pair{op.num + 1, v.right}
				r += follow(op.next, cc)
				values[op.desc] = pair{v.left, op.num}
			}
		}
		return r
	}
	in := map[byte]pair{
		'x': pair{1, 4000},
		'm': pair{1, 4000},
		'a': pair{1, 4000},
		's': pair{1, 4000},
	}
	p2 := follow("in", in)
	fmt.Println("Part 2:", p2)
}

func clone(data map[byte]pair) map[byte]pair {
	r := map[byte]pair{}
	for k, v := range data {
		r[k] = v
	}
	return r
}
