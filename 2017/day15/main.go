package main

import "fmt"

var (
	divisor = uint64(2147483647)
	compare = uint64(0xffff)
)

type Generator struct {
	Start   uint64
	Factor  uint64
	Current uint64
}

func Count(iterations int, A, B Generator) int {
	A.Current = A.Start
	B.Current = B.Start
	count := 0
	for i := 0; i < iterations; i++ {
		A.Current = (A.Current * A.Factor) % divisor
		for A.Current%4 != 0 {
			A.Current = (A.Current * A.Factor) % divisor
		}
		B.Current = (B.Current * B.Factor) % divisor
		for B.Current%8 != 0 {
			B.Current = (B.Current * B.Factor) % divisor
		}
		if A.Current&compare == B.Current&compare {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println("Test: ", Count(1057,
		Generator{Start: 65, Factor: 16807},
		Generator{Start: 8921, Factor: 48271},
	))
	fmt.Println("Final: ", Count(5000000,
		Generator{Start: 277, Factor: 16807},
		Generator{Start: 349, Factor: 48271},
	))
}
