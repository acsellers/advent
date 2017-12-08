package main

import (
	"fmt"
	"strings"
)

func Dist(path string) int {
	dirs := strings.Split(path, ", ")
	loc := [2]int{0, 0}
	northSouth := true
	positive := true
	/*
		C Cns Cp|R Rns Rp|L Lns Lp
		N T   T |E F   T |W F   F
		E F   T |S T   F |N T   T
		S T   F |W       |E
		W F   F |N       |S T   F
	*/
	for _, dir := range dirs {

	}
	return len(dirs)
}
func main() {
	fmt.Println('1' - '0')
}
