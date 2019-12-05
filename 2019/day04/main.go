package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	for i := 128392; i <= 643281; i++ {
		if Check(i) {
			fmt.Println(i)
		}
	}
}

func Check(i int) bool {
	r := []rune(fmt.Sprint(i))
	last := r[0]
	repeat := false
	ascend := true
	for i, ir := range r {
		if ir == last && i > 0 {
			repeat = true
		}

		li, err := strconv.Atoi(string([]rune{last}))
		if err != nil {
			log.Fatal(err)
		}
		ci, err := strconv.Atoi(string([]rune{ir}))
		if err != nil {
			log.Fatal(err)
		}
		if ci < li {
			ascend = false
		}
		last = ir
	}
	if repeat && ascend {
		dupes := map[rune]int{}
		for i := 1; i < len(r); i++ {
			if r[i-1] == r[i] {
				dupes[r[i]]++

			}
		}
		for _, v := range dupes {
			if v == 1 {
				return true
			}
		}

	}
	return false
}
