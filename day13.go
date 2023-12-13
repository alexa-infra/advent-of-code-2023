package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	patterns := [][]string{}
	pattern := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			patterns = append(patterns, pattern)
			pattern = []string{}
		} else {
			pattern = append(pattern, line)
		}
	}
	patterns = append(patterns, pattern)

	findReflectionsRow := func(pat []string, err int) (bool, int, int) {
		n, m := len(pat), len(pat[0])
		max, maxRow := 0, 0
		for i := 1; i < n; i++ {
			smudge := err
			rowMatch := true
			for j := 0; j < m; j++ {
				if pat[i][j] != pat[i-1][j] {
					if smudge == 0 {
						rowMatch = false
						break
					} else {
						smudge--
					}
				}
			}
			if rowMatch {
				mirror := true
				k := 1
				for i+k < n && i-1-k >= 0 {
					for j := 0; j < m; j++ {
						if pat[i+k][j] != pat[i-1-k][j] {
							if smudge == 0 {
								mirror = false
								break
							} else {
								smudge--
							}
						}
						if !mirror {
							break
						}
					}
					k++
				}
				if mirror && smudge == 0 {
					if 2+k-1 > max {
						max = 2 + k - 1
						maxRow = i
					}
				}
			}
		}
		return max > 0, maxRow, max * m
	}
	findReflectionsCol := func(pat []string, err int) (bool, int, int) {
		n, m := len(pat), len(pat[0])
		max, maxCol := 0, 0
		for i := 1; i < m; i++ {
			smudge := err
			colMatch := true
			for j := 0; j < n; j++ {
				if pat[j][i] != pat[j][i-1] {
					if smudge == 0 {
						colMatch = false
						break
					} else {
						smudge--
					}
				}
			}
			if colMatch {
				mirror := true
				k := 1
				for i+k < m && i-1-k >= 0 {
					for j := 0; j < n; j++ {
						if pat[j][i+k] != pat[j][i-1-k] {
							if smudge == 0 {
								mirror = false
								break
							} else {
								smudge--
							}
						}
					}
					if !mirror {
						break
					}
					k++
				}
				if mirror && smudge == 0 {
					if 2+k-1 > max {
						max = 2 + k - 1
						maxCol = i
					}
				}
			}
		}
		return max > 0, maxCol, max * n
	}
	findReflections := func(err int) int {
		r := 0
		for index, pat := range patterns {
			found1, col, area1 := findReflectionsCol(pat, err)
			found2, row, area2 := findReflectionsRow(pat, err)
			if found1 && found2 {
				if area1 > area2 {
					r += col
				} else {
					r += row * 100
				}
			} else if found1 {
				r += col
			} else if found2 {
				r += row * 100
			} else {
				fmt.Println(index, "Not found")
			}
		}
		return r
	}
	p1 := findReflections(0)
	fmt.Println("Part 1:", p1)
	p2 := findReflections(1)
	fmt.Println("Part 2:", p2)
}
