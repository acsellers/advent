package main

import (
	"fmt"
	"math/bits"
)

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

func Count(key string) int {
	used := 0
	for i := 0; i < 128; i++ {
		h := Hash([]uint8([]byte(fmt.Sprintf("%s-%d", key, i))))
		for _, u := range h {

			used += bits.OnesCount8(u)
		}
	}
	return used
}
func Regions(key string) uint16 {
	regions := [][]uint16{}
	var max uint16
	max--
	for i := 0; i < 128; i++ {
		row := make([]uint16, 128)
		h := Hash([]uint8([]byte(fmt.Sprintf("%s-%d", key, i))))
		for i, u := range h {
			if u&0x80 > 0 {
				row[i*8] = max
			}
			if u&0x40 > 0 {
				row[i*8+1] = max
			}
			if u&0x20 > 0 {
				row[i*8+2] = max
			}
			if u&0x10 > 0 {
				row[i*8+3] = max
			}
			if u&0x08 > 0 {
				row[i*8+4] = max
			}
			if u&0x04 > 0 {
				row[i*8+5] = max
			}
			if u&0x02 > 0 {
				row[i*8+6] = max
			}
			if u&0x01 > 0 {
				row[i*8+7] = max
			}
		}
		regions = append(regions, row)
	}
	var maxGroup uint16
	at := func(x, y int) uint16 {
		if x < 0 || x > 127 {
			return 0
		}
		if y < 0 || y > 127 {
			return 0
		}
		return regions[x][y]
	}
	var color func(x, y int)
	color = func(x, y int) {
		if at(x, y) != max {
			return
		}
		regions[x][y] = maxGroup
		if at(x-1, y) == max {
			color(x-1, y)
		}
		if at(x+1, y) == max {
			color(x+1, y)
		}
		if at(x, y-1) == max {
			color(x, y-1)
		}
		if at(x, y+1) == max {
			color(x, y+1)
		}
	}
	for x, row := range regions {
		for y, cell := range row {
			if cell == max {
				maxGroup++
				color(x, y)
			}
		}
	}
	return maxGroup
}
func main() {
	fmt.Println("Test: ", Count("flqrgnkx"))
	fmt.Println("Input: ", Count("jxqlasbh"))
	fmt.Println("Test: ", Regions("flqrgnkx"))
	fmt.Println("Input: ", Regions("jxqlasbh"))
}
