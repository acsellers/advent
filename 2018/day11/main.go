package main

import "fmt"

func main() {
	fmt.Println("4 ==", Calc(8, 3, 5))
	fmt.Println("-5 ==", Calc(57, 122, 79))
	fmt.Println("0 ==", Calc(39, 217, 196))
	fmt.Println("4 ==", Calc(71, 101, 153))
	fmt.Println(Do(18))
	fmt.Println(Do(42))
	fmt.Println(Do(7672))
	fmt.Println(Do2(18))
	fmt.Println(Do2(42))
	fmt.Println(Do2(7672))
}

func Do(serial int) (int, int) {
	grid := [300][300]int{}
	for row := range grid {
		for col := range grid[row] {
			grid[row][col] = Calc(serial, row, col)
		}
	}
	/*for _, row := range grid {
		fmt.Println(row)
	}*/
	max, maxX, maxY := 0, 0, 0
	for row := range grid {
		if row+2 < len(grid) {
			for col := range grid[row] {
				if col+2 < len(grid) {
					sum := grid[row][col] + grid[row+1][col] + grid[row+2][col]
					sum += grid[row][col+1] + grid[row+1][col+1] + grid[row+2][col+1]
					sum += grid[row][col+2] + grid[row+1][col+2] + grid[row+2][col+2]
					if sum > max {
						max = sum
						maxX = row
						maxY = col
					}
				}
			}
		}
	}
	return maxX, maxY
}
func Calc(serial, x, y int) int {
	rackid := x + 10
	return ((((rackid*y)+serial)*rackid)/100)%10 - 5
}

func Do2(serial int) (int, int, int, int) {
	grid := [300][300]int{}
	for row := range grid {
		for col := range grid[row] {
			grid[row][col] = Calc(serial, row, col)
		}
	}
	/*for _, row := range grid {
		fmt.Println(row)
	}*/
	size, max, maxX, maxY := 0, 0, 0, 0
	for i := 1; i < 32; i++ {
		for row := range grid {
			if row+i <= len(grid) {
				for col := range grid[row] {
					if col+i <= len(grid) {
						sum := 0
						for j := 0; j < i; j++ {
							for k := 0; k < i; k++ {
								sum += grid[row+j][col+k]
							}
						}

						if sum > max {
							max = sum
							maxX = row
							maxY = col
							size = i
						}
					}
				}
			}
		}
	}
	return maxX, maxY, size, max
}
