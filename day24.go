package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Particle struct {
	x, y, z    int
	vx, vy, vz int
}

func main() {
	particles := []Particle{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " @ ")
		pos := strings.Split(parts[0], ", ")
		vel := strings.Split(parts[1], ", ")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		z, _ := strconv.Atoi(pos[2])
		vx, _ := strconv.Atoi(vel[0])
		vy, _ := strconv.Atoi(vel[1])
		vz, _ := strconv.Atoi(vel[2])
		particles = append(particles, Particle{x, y, z, vx, vy, vz})
	}
	n := len(particles)
	p1 := 0
	for i, a := range particles {
		for j := i + 1; j < n; j++ {
			b := particles[j]
			d := a.vy*b.vx - a.vx*b.vy
			if d == 0 {
				continue
			}
			t2 := float64((b.y-a.y)*a.vx-(b.x-a.x)*a.vy) / float64(d)
			if t2 < 0 {
				continue
			}
			if a.vx == 0 {
				continue
			}
			t1 := (float64(b.x-a.x) + float64(b.vx)*t2) / float64(a.vx)
			if t1 < 0 {
				continue
			}
			x := float64(a.x) + float64(a.vx)*t1
			y := float64(a.y) + float64(a.vy)*t1
			if x >= 200000000000000 && x <= 400000000000000 && y >= 200000000000000 && y <= 400000000000000 {
				p1++
			}
		}
	}
	fmt.Println("Part 1:", p1)
	//a + av * t1 = stone.p + stone.v * t1 // 7 var, 3 eq
	//b + bv * t2 = stone.p + stone.v * t2 // 8 var, 6 eq
	//c + cv * t3 = stone.p + stone.v * t3 // 9 var, 9 eq

	a, b, c := particles[0], particles[1], particles[2]

	//a.x + a.vx * t1 = stone.x + stone.vx * t1
	//a.y + a.vy * t1 = stone.y + stone.vy * t1
	//a.z + a.vz * t1 = stone.z + stone.vz * t1
	//b.x + b.vx * t2 = stone.x + stone.vx * t2
	//b.y + b.vy * t2 = stone.y + stone.vy * t2
	//b.z + b.vz * t2 = stone.z + stone.vz * t2
	//c.x + c.vx * t3 = stone.x + stone.vx * t3
	//c.y + c.vy * t3 = stone.y + stone.vy * t3
	//c.z + c.vz * t3 = stone.z + stone.vz * t3

	//(vx - avx) * t1                                 + sx           = ax
	//(vy - avy) * t1                                      + sy      = ay
	//(vz - avz) * t1                                           + sz = az
	//               (vx - bvx) * t2                  + sx           = bx
	//               (vy - bvy) * t2                       + sy      = by
	//               (vz - bvz) * t2                            + sz = bz
	//                                (vx - cvx) * t3 + sx           = cx
	//                                (vy - cvy) * t3      + sy      = cy
	//                                (vz - cvz) * t3           + sz = cz
	//
	//	(vx-avx) * t1 + (bvx - vx) * t2 = ax - bx -> t2 = (ax - bx - (vx-avx) * t1) / (bvx-vx)
	//	(vy-avy) * t1 + (bvy - vy) * t2 = ay - by -> t1 = (ay - by - (bvy -vy) * t2) / (vy-avy)
	//
	//(vx-avx) * ((ay - by) - (bvy -vy) * t2) + (vy - avy) * (bvx - vx) * t2 = (ax - bx) * (vy -avy)
	//(1-avx) * (ay - by) - (1-avx) * (bvy -1) * t2 + (1 - avy) * (bvx - 1) * t2 = (ax - bx) * (1 -avy)
	//((1 - avy) * (bvx - 1) - (1-avx) * (bvy -1)) * t2 = (ax - bx) * (1 -avy) - (1-avx) * (ay-by)
	//t2 = ((ax-bx) * (1-avy) - (1-avx) * (ay-by)) / ((1 - avy) * (bvx - 1) - (1-avx) * (bvy -1))

	boundary := 300
	for vx := -boundary; vx <= boundary; vx++ {
		for vy := -boundary; vy <= boundary; vy++ {
			if ((vy-a.vy)*(b.vx-vx) - (vx-a.vx)*(b.vy-vy)) == 0 {
				continue
			}
			t2 := ((a.x-b.x)*(vy-a.vy) - (vx-a.vx)*(a.y-b.y)) / ((vy-a.vy)*(b.vx-vx) - (vx-a.vx)*(b.vy-vy))
			if t2 < 0 || vy == a.vy {
				continue
			}
			t1 := (a.y - b.y - (b.vy-vy)*t2) / (vy - a.vy)
			if t1 < 0 {
				continue
			}
			for vz := -boundary; vz <= boundary; vz++ {
				sx := a.x - (vx-a.vx)*t1
				sy := a.y - (vy-a.vy)*t1
				sz := a.z - (vz-a.vz)*t1
				if vz == c.vz {
					continue
				}
				t3 := (c.z - sz) / (vz - c.vz)
				if t3 < 0 {
					continue
				}
				sx1 := b.x - (vx-b.vx)*t2
				sy1 := b.y - (vy-b.vy)*t2
				sz1 := b.z - (vz-b.vz)*t2
				if sx != sx1 || sy != sy1 || sz != sz1 {
					continue
				}
				sx2 := c.x - (vx-c.vx)*t3
				sy2 := c.y - (vy-c.vy)*t3
				sz2 := c.z - (vz-c.vz)*t3
				if sx != sx2 || sy != sy2 || sz != sz2 {
					continue
				}
				fmt.Println("Part 2:", sx+sy+sz)
				return
			}
		}
	}
}
