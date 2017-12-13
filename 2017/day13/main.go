package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Firewall struct {
	Layers map[int]*Layer
	Max    int
}
type Layer struct {
	Height  int
	Current int
	Reverse bool
}

func (f *Firewall) Step() {
	for _, l := range f.Layers {
		if l.Reverse {
			if l.Current == 0 {
				l.Reverse = false
				l.Current++
			} else {
				l.Current--
			}
		} else {
			l.Current++
			if l.Current == l.Height {
				l.Reverse = true
				l.Current -= 2
			}
		}
	}
}

func (f Firewall) Caught(layer int) int {
	if l, ok := f.Layers[layer]; ok {
		if l.Current == 0 {
			return layer * l.Height
		}
	}
	return 0
}

func Severity(input string) int {
	f := Firewall{Layers: make(map[int]*Layer)}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var layer, height int
		fmt.Sscanf(scanner.Text(), "%d: %d", &layer, &height)
		if layer > f.Max {
			f.Max = layer + 1
		}
		f.Layers[layer] = &Layer{Height: height}
	}
	severity := 0
	for i := 0; i <= f.Max; i++ {
		severity += f.Caught(i)
		f.Step()
	}
	return severity
}

type Eq struct {
	M, B int
}

func (e Eq) String() string {
	return fmt.Sprintf("{M: %d, B: %d}\n", e.M, e.B)
}

func Invisible(input string, incr int) int {
	eqs := []Eq{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var layer, height int
		fmt.Sscanf(scanner.Text(), "%d: %d", &layer, &height)

		eqs = append(eqs, Eq{M: 2*height - 2, B: 2*(height-1) - layer})
	}
	i := 0
Loop:
	for {
		for _, eq := range eqs {
			if (i-eq.B)%eq.M == 0 {
				i++
				continue Loop
			}
		}

		return i
	}
}

func main() {
	fmt.Println("Test1: ", Severity("0: 3\n1: 2\n4: 4\n6: 4"))
	fmt.Println("Part1: ", Severity(input))
	fmt.Println("Test2: ", Invisible("0: 3\n1: 2\n4: 4\n6: 4", 1))
	fmt.Println("Part2: ", Invisible(input, 2))
}

var input = `0: 4
1: 2
2: 3
4: 4
6: 8
8: 5
10: 8
12: 6
14: 6
16: 8
18: 6
20: 6
22: 12
24: 12
26: 10
28: 8
30: 12
32: 8
34: 12
36: 9
38: 12
40: 8
42: 12
44: 17
46: 14
48: 12
50: 10
52: 20
54: 12
56: 14
58: 14
60: 14
62: 12
64: 14
66: 14
68: 14
70: 14
72: 12
74: 14
76: 14
80: 14
84: 18
88: 14`
