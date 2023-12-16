package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	n, m := len(lines), len(lines[0])
	type pos struct {
		x, y int
	}
	type particle struct {
		p   pos
		dir byte
	}
	calc := func(start particle) int {
		visited := map[pos]byte{}
		lights := []particle{ start }
		for len(lights) > 0 {
			light := lights[0]
			lights = lights[1:]
			if light.p.x < 0 || light.p.x >= n || light.p.y < 0 || light.p.y >= m {
				continue
			}
			val, ok := visited[light.p]
			if ok && val == light.dir {
				continue
			}
			visited[light.p] = light.dir
			ch := lines[light.p.x][light.p.y]
			if ch == '.' {
				next := light.p
				if light.dir == 'E' {
					next.y += 1
				} else if light.dir == 'S' {
					next.x += 1
				} else if light.dir == 'W' {
					next.y -= 1
				} else if light.dir == 'N' {
					next.x -= 1
				}
				lights = append(lights, particle{next, light.dir})
			} else if ch == '/' {
				next := light.p
				dir := light.dir
				if light.dir == 'E' {
					next.x -= 1
					dir = 'N'
				} else if light.dir == 'S' {
					next.y -= 1
					dir = 'W'
				} else if light.dir == 'W' {
					next.x += 1
					dir = 'S'
				} else if light.dir == 'N' {
					next.y += 1
					dir = 'E'
				}
				lights = append(lights, particle{next, dir})
			} else if ch == '\\' {
				next := light.p
				dir := light.dir
				if light.dir == 'E' {
					next.x += 1
					dir = 'S'
				} else if light.dir == 'S' {
					next.y += 1
					dir = 'E'
				} else if light.dir == 'W' {
					next.x -= 1
					dir = 'N'
				} else if light.dir == 'N' {
					next.y -= 1
					dir = 'W'
				}
				lights = append(lights, particle{next, dir})
			} else if ch == '|' {
				next := light.p
				next2 := light.p
				if light.dir == 'E' || light.dir == 'W' {
					next.x += 1
					next2.x -= 1
					lights = append(lights, particle{next, 'S'})
					lights = append(lights, particle{next2, 'N'})
				} else if light.dir == 'S' {
					next.x += 1
					lights = append(lights, particle{next, light.dir})
				} else if light.dir == 'N' {
					next.x -= 1
					lights = append(lights, particle{next, light.dir})
				}
			} else if ch == '-' {
				next := light.p
				next2 := light.p
				if light.dir == 'E' {
					next.y += 1
					lights = append(lights, particle{next, light.dir})
				} else if light.dir == 'W' {
					next.y -= 1
					lights = append(lights, particle{next, light.dir})
				} else if light.dir == 'N' || light.dir == 'S' {
					next.y -= 1
					next2.y += 1
					lights = append(lights, particle{next, 'W'})
					lights = append(lights, particle{next2, 'E'})
				}
			}
		}
		return len(visited)
	}
	p1 := calc(particle{pos{0, 0}, 'E'})
	fmt.Println("Part 1:", p1)
	p2 := 0
	starts := []particle {
		particle{pos{0, 0}, 'E'},
		particle{pos{0, 0}, 'S'},
		particle{pos{n - 1, 0}, 'N'},
		particle{pos{n - 1, 0}, 'E'},
		particle{pos{0, m - 1}, 'W'},
		particle{pos{0, m - 1}, 'S'},
		particle{pos{n - 1, m - 1}, 'W'},
		particle{pos{n - 1, m - 1}, 'N'},
	}
	for i := 1; i < n - 1; i++ {
		starts = append(starts, particle{pos{i, 0}, 'W'})
		starts = append(starts, particle{pos{i, m - 1}, 'E'})
	}
	for i := 1; i < m - 1; i++ {
		starts = append(starts, particle{pos{0, i}, 'S'})
		starts = append(starts, particle{pos{n - 1, i}, 'N'})
	}
	for _, s := range starts {
		r := calc(s)
		if r > p2 {
			p2 = r
		}
	}
	fmt.Println("Part 2:", p2)
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		p := pos{i, j}
	//		_, ok := visited[p]
	//		if ok {
	//			fmt.Print("#")
	//		} else {
	//			fmt.Print(".")
	//		}
	//	}
	//	fmt.Println()
	//}
}
