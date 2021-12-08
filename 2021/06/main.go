package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Do(example)
	Do2(example)
	Do(input)
	Do2(input)
}
func Do(input string) {
	fishs := strings.Split(input, ",")
	fishes := []int{}
	for _, fish := range fishs {
		i, _ := strconv.Atoi(fish)
		fishes = append(fishes, i)
	}
	for i := 0; i < 80; i++ {
		next := []int{}
		babies := []int{}
		for i := range fishes {
			if fishes[i] == 0 {
				next = append(next, 6)
				babies = append(babies, 8)
			} else {
				next = append(next, fishes[i]-1)
			}
		}

		fishes = append(next, babies...)
	}
	fmt.Println(len(fishes))

}
func Do2(input string) {
	counts := make([]int, 9)
	fishs := strings.Split(input, ",")
	for _, fish := range fishs {
		i, _ := strconv.Atoi(fish)
		counts[i]++
	}
	for i := 0; i < 256; i++ {
		next := append(counts[1:], 0)
		next[6] += counts[0]
		next[8] += counts[0]
		counts = next
	}
	sum := 0
	for i := range counts {
		sum += counts[i]
	}
	fmt.Println(sum)
}

var example = `3,4,3,1,2`

var input = `4,3,3,5,4,1,2,1,3,1,1,1,1,1,2,4,1,3,3,1,1,1,1,2,3,1,1,1,4,1,1,2,1,2,2,1,1,1,1,1,5,1,1,2,1,1,1,1,1,1,1,1,1,3,1,1,1,1,1,1,1,1,5,1,4,2,1,1,2,1,3,1,1,2,2,1,1,1,1,1,1,1,1,1,1,4,1,3,2,2,3,1,1,1,4,1,1,1,1,5,1,1,1,5,1,1,3,1,1,2,4,1,1,3,2,4,1,1,1,1,1,5,5,1,1,1,1,1,1,4,1,1,1,3,2,1,1,5,1,1,1,1,1,1,1,5,4,1,5,1,3,4,1,1,1,1,2,1,2,1,1,1,2,2,1,2,3,5,1,1,1,1,3,5,1,1,1,2,1,1,4,1,1,5,1,4,1,2,1,3,1,5,1,4,3,1,3,2,1,1,1,2,2,1,1,1,1,4,5,1,1,1,1,1,3,1,3,4,1,1,4,1,1,3,1,3,1,1,4,5,4,3,2,5,1,1,1,1,1,1,2,1,5,2,5,3,1,1,1,1,1,3,1,1,1,1,5,1,2,1,2,1,1,1,1,2,1,1,1,1,1,1,1,3,3,1,1,5,1,3,5,5,1,1,1,2,1,2,1,5,1,1,1,1,2,1,1,1,2,1`
