package main

import (
	"fmt"
	"sort"
)

func Who(count int) int {
	remaining := map[int]bool{}
	indexes := []int{}
	for i := 1; i <= count; i++ {
		remaining[i] = true
		indexes = append(indexes, i)
	}
	for len(remaining) > 1 {
	CheckLoop:
		for _, i := range indexes {
			if remaining[i] {
				for j := i + 1; j <= count; j++ {
					if remaining[j] {
						delete(remaining, j)
						continue CheckLoop
					}
				}
				for j := 1; j < i; j++ {
					if remaining[j] {
						delete(remaining, j)
						continue CheckLoop
					}
				}
			}
		}
		indexes = make([]int, 0, len(remaining))
		for k := range remaining {
			indexes = append(indexes, k)
		}
		sort.Ints(indexes)

	}
	for k := range remaining {
		return k
	}
	return 0
}

func Who2(count int) int {
	indexes := []int{}
	for i := 1; i <= count; i++ {
		indexes = append(indexes, i)
	}
	current := 0
	for len(indexes) > 1 {
		trade := current + len(indexes)/2

		//if len(indexes)%2 == 0 {
		//	trade++
		//}
		if trade >= len(indexes) {
			trade -= len(indexes)
		}
		indexes = append(indexes[:trade], indexes[trade+1:]...)
		if trade < current {
			current--
		} else {
			current++
		}

		fmt.Println(indexes)
		//fmt.Println("Len: ", len(indexes))
	}
	//fmt.Println(indexes)
	return indexes[0]
}
func main() {
	fmt.Println("5: ", Who(5))
	//fmt.Println("Input: ", Who(3018458))
	fmt.Println("5: ", Who2(5))
	fmt.Println("5: ", Who2(14))
	//fmt.Println("Input: ", Who2(3018458))
}
