package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	timeLine := scanner.Text()
	timeMatch := re.FindAllString(timeLine, -1)
	scanner.Scan()
	distanceLine := scanner.Text()
	distanceMatch := re.FindAllString(distanceLine, -1)
	n := len(timeMatch)
	p1 := 1
	for i := 0; i < n; i++ {
		time, _ := strconv.Atoi(timeMatch[i])
		dist, _ := strconv.Atoi(distanceMatch[i])
		cc := 0
		//t1 + t2 = time
		//t1 * t2 = dist
		for t1 := 0; t1 <= time; t1++ {
			t2 := time - t1
			if t1*t2 > dist {
				cc++
			}
		}
		p1 *= cc
	}
	fmt.Println("Part 1:", p1)
	timeStr := ""
	distStr := ""
	for i := 0; i < n; i++ {
		timeStr += timeMatch[i]
		distStr += distanceMatch[i]
	}
	time, _ := strconv.Atoi(timeStr)
	dist, _ := strconv.Atoi(distStr)
	//t1 * (time - t1) = dist
	//t1 * t2 - t1*t1 = dist
	// t1*t1 - t1*time + dist = 0
	// a x*x + b x + c = 0
	// D = b*b - 4ac
	D := float64(time*time - 4*dist)
	// (-b +/- sqrt(D)) / 2a
	sqrtD := math.Sqrt(D)
	x1 := math.Ceil((float64(time) - sqrtD) / 2)
	x2 := math.Floor((float64(time) + sqrtD) / 2)
	p2 := int64(x2) - int64(x1) + 1
	fmt.Println("Part 2:", p2)
}
