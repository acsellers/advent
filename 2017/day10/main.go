package main

import "fmt"

func flip(v []uint8) []uint8 {
	for i := 0; i < len(v)/2; i++ {
		t := v[i]
		v[i] = v[len(v)-1-i]
		v[len(v)-1-i] = t
	}
	return v
}
func Encrypt(vals, swaps []uint8, ix, swp int) ([]uint8, int, int) {
	count := len(vals)
	read := func(index int, length uint8) []uint8 {
		rv := make([]uint8, length)
		for i := range rv {
			if index == count {
				index = 0
			}
			rv[i] = vals[index]
			index++
		}
		return rv
	}
	write := func(ix int, rv []uint8) {
		for _, v := range rv {
			if ix == count {
				ix = 0
			}
			vals[ix] = v
			ix++
		}
	}
	for _, sw := range swaps {
		tv := read(ix, sw)
		tv = flip(tv)
		write(ix, flip(read(ix, sw)))
		ix = (ix + swp + int(sw)) % count
		swp++
	}
	return vals, ix, swp
}

func Hash(input []uint8) []uint8 {
	vals := make([]uint8, 256)
	for i := range vals {
		vals[i] = uint8(i)
	}
	input = append(input, 17, 31, 73, 47, 23)
	ix, swp := 0, 0
	for i := 0; i < 64; i++ {
		vals, ix, swp = Encrypt(vals, input, ix, swp)
		//fmt.Println("Round: ", vals)
	}
	out := make([]uint8, 16)
	for i := 0; i < 256; i += 16 {
		out[i/16] = vals[i] ^ vals[i+1] ^ vals[i+2] ^ vals[i+3] ^ vals[i+4] ^ vals[i+5] ^ vals[i+6] ^ vals[i+7] ^ vals[i+8] ^ vals[i+9] ^ vals[i+10] ^ vals[i+11] ^ vals[i+12] ^ vals[i+13] ^ vals[i+14] ^ vals[i+15]
	}
	return out
}

func main() {

	fmt.Println(Encrypt(
		[]uint8{0, 1, 2, 3, 4},
		[]uint8{3, 4, 1, 5},
		0, 0,
	))
	vals := make([]uint8, 256)
	for i := range vals {
		vals[i] = uint8(i)
	}
	vals, _, _ = Encrypt(vals, []uint8{34, 88, 2, 222, 254, 93, 150, 0, 199, 255, 39, 32, 137, 136, 1, 167}, 0, 0)
	fmt.Println("Blah ", vals)
	fmt.Println(int(vals[0]) * int(vals[1]))
	fmt.Printf("%x\n", Hash([]uint8{}))
	fmt.Printf("%x\n", Hash([]uint8([]byte("AoC 2017"))))
	fmt.Printf("%x\n", Hash([]uint8([]byte("1,2,3"))))
	fmt.Printf("%x\n", Hash([]uint8([]byte("1,2,4"))))
	fmt.Printf("%x\n", Hash([]uint8([]byte("34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167"))))
}
