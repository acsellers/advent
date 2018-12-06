package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Do(demo))
	fmt.Println(Do(input))
	fmt.Println(Do2(input))
}

func Do(input string) int {
	maxX, maxY := 0, 0
	points := []Point{}
	for _, pt := range strings.Split(input, "\n") {
		parts := strings.Split(pt, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		points = append(points, Point{x + 5, y + 5})
	}

	board := make([][]int, (maxX + 10))
	for i := 0; i < maxX+10; i++ {
		board[i] = make([]int, maxY+10)
	}

	for i := 0; i < maxX+10; i++ {
		for j := 0; j < maxY+10; j++ {
			dist := 999999
			pt1, pt2 := -1, -1
			for k, pt := range points {
				d := pt.Distance(i, j)
				if d < dist {
					dist = d
					dist = d
					pt1 = k
					pt2 = k
				}
				if d == dist {
					pt2 = k
				}
			}
			if pt1 == pt2 {
				board[i][j] = pt1
			} else {
				board[i][j] = -1
			}
		}
	}
	totals := make([]int, len(points))
	for i := 0; i < maxX+10; i++ {
		for j := 0; j < maxY+10; j++ {
			if board[i][j] != -1 {
				totals[board[i][j]]++
			}
		}
	}
	totals2 := make([]int, len(points))
	for i := 2; i < maxX+8; i++ {
		for j := 2; j < maxY+8; j++ {
			if board[i][j] != -1 {
				totals2[board[i][j]]++
			}
		}
	}
	max := 0
	for i, t := range totals {
		if t > max && totals2[i] == t {
			max = t
		}
	}
	return max
}

type Point struct {
	X, Y int
}

func (p Point) Distance(x, y int) int {
	if x > p.X {
		if y > p.Y {
			return x - p.X + y - p.Y
		} else {
			return x - p.X + p.Y - y
		}
	} else {
		if y > p.Y {
			return p.X - x + y - p.Y
		} else {
			return p.X - x + p.Y - y
		}
	}
}

var demo = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

var input = `224, 153
176, 350
353, 241
207, 59
145, 203
123, 210
113, 203
191, 241
172, 196
209, 249
260, 229
98, 231
305, 215
258, 141
337, 282
156, 140
325, 197
179, 279
283, 233
317, 150
305, 245
67, 109
251, 140
245, 59
173, 105
59, 173
257, 70
269, 110
102, 162
179, 180
324, 112
357, 311
317, 245
239, 112
321, 220
133, 97
334, 99
117, 102
133, 112
222, 316
68, 296
150, 287
263, 263
66, 347
128, 118
63, 202
68, 236
264, 122
77, 243
92, 110`

func Do2(input string) int {
	maxX, maxY := 0, 0
	points := []Point{}
	for _, pt := range strings.Split(input, "\n") {
		parts := strings.Split(pt, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		points = append(points, Point{x + 5, y + 5})
	}

	board := make([][]int, (maxX + 10))
	for i := 0; i < maxX+10; i++ {
		board[i] = make([]int, maxY+10)
	}

	cnt := 0
	for i := 0; i < maxX+10; i++ {
		for j := 0; j < maxY+10; j++ {
			dist := 0
			for _, pt := range points {
				dist += pt.Distance(i, j)
			}
			if dist < 10000 {
				cnt++
				if i == 0 || j == 0 {
					log.Fatal("Bigger space needed")
				}
			}
		}
	}
	return cnt
}
