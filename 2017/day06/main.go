package main

import "fmt"

func maxIndex(is []int) int {
	ix := 0
	m := is[0]
	for i := 1; i < len(is); i++ {
		if is[i] > m {
			m = is[i]
			ix = i
		}
	}
	return ix
}

func Distribute(banks []int) int {
	seen := make(map[string]bool)
	seen[fmt.Sprint(banks)] = true
	iterations := 0
	for {
		iterations++
		ix := maxIndex(banks)
		temp := banks[ix]
		banks[ix] = 0
		ix++
		for temp > 0 {
			if ix == len(banks) {
				ix = 0
			}
			banks[ix]++
			ix++
			temp--
		}

		cfg := fmt.Sprint(banks)
		if seen[cfg] {
			return iterations
		}
		seen[cfg] = true
	}

}
func Distribute2(banks []int) int {
	seen := make(map[string]bool)
	seen[fmt.Sprint(banks)] = true
	var first string
	for first == "" {
		ix := maxIndex(banks)
		temp := banks[ix]
		banks[ix] = 0
		ix++
		for temp > 0 {
			if ix == len(banks) {
				ix = 0
			}
			banks[ix]++
			ix++
			temp--
		}

		cfg := fmt.Sprint(banks)
		if seen[cfg] {
			first = cfg
		}
		seen[cfg] = true
	}
	iterations := 0
	for {
		iterations++
		ix := maxIndex(banks)
		temp := banks[ix]
		banks[ix] = 0
		ix++
		for temp > 0 {
			if ix == len(banks) {
				ix = 0
			}
			banks[ix]++
			ix++
			temp--
		}

		cfg := fmt.Sprint(banks)
		if first == cfg {
			return iterations
		}
	}
}
func main() {
	fmt.Println("Dist1 Test: ", Distribute([]int{0, 2, 7, 0}))
	input := []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}
	fmt.Println("Dist1: ", Distribute(input))
	fmt.Println("Dist2 Test: ", Distribute2([]int{0, 2, 7, 0}))
	fmt.Println("Dist2: ", Distribute2(input))

}
