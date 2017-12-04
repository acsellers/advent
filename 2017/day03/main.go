package main

import "fmt"

var ends = make([]int, 500)

func init() {
	ends[0] = 1
	ends[1] = 9
	ends[2] = 25

	for i := 1; i < 500; i++ {
		begin := ends[i-1]
		length := i*2 + 1
		ends[i] = begin + (length-1)*4
		//fmt.Println(starts[i])
	}
	fmt.Println(ends)
}
func Solve(input int) int {
	var ring, offset, width int
	for i, v := range ends {
		if input <= v {
			ring = i
			width = (ends[i] - ends[i-1]) / 4
			offset = input - ends[i-1]
			break
		}
	}
	for width < offset {
		offset = offset - width
	}

	return ring + offset - ring
}

var table [][]int

func init() {

}
func At(x, y int) int {
	return table[x][y]
}

func Solve2(input int) int {
	table = make([][]int, 0, 21)
	for i := 0; i < 22; i++ {
		table = append(table, make([]int, 21))
	}
	table[11][11] = 1
	ring := 2
	x := 11
	y := 11
	for j := 0; j < 4; j++ {
		// move one ring out
		y++
		table[x][y] = At(x, y-1) + At(x-1, y-1)
		if table[x][y] > input {
			return table[x][y]
		}
		// move up
		for i := 1; i < ring; i++ {
			x--
			table[x][y] = At(x+1, y-1) + At(x+1, y) + At(x-1, y-1) + At(x, y-1)
			if table[x][y] > input {
				return table[x][y]
			}
		}
		// move left
		for i := 0; i < ring; i++ {
			y--
			table[x][y] = At(x, y+1) + At(x+1, y-1) + At(x+1, y) + At(x+1, y+1)
			if table[x][y] > input {
				return table[x][y]
			}
		}
		// move down
		for i := 0; i < ring; i++ {
			x++
			table[x][y] = At(x-1, y) + At(x-1, y+1) + At(x, y+1) + At(x+1, y+1)
			if table[x][y] > input {
				return table[x][y]
			}
		}
		// move right
		for i := 0; i < ring; i++ {
			y++

			table[x][y] = At(x, y-1) + At(x-1, y-1) + At(x-1, y) + At(x-1, y+1)
			if table[x][y] > input {
				return table[x][y]
			}
		}

		ring = ring + 2
	}
	return 0
}

func main() {
	fmt.Println("Solve 1: ", Solve(12))
	fmt.Println("Solve 1: ", Solve(23))
	fmt.Println("Solve 1: ", Solve(16))
	fmt.Println("Solve 1: ", Solve(1024))
	fmt.Println("Solve 1: ", Solve(347991))
	fmt.Println("Solve 2: ", Solve2(23))
	fmt.Println("Solve 2: ", Solve2(55))
	fmt.Println("Solve 2: ", Solve2(347991))

}
