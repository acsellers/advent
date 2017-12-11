package main

import (
	"bufio"
	"strings"
)

type Board struct {
	Pixels [][]bool
}

func (b *Board) Execute(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		switch {
		case strings.HasPrefix(scanner.Text(), "rect"):
			
		case strings.HasPrefix(scanner.Text(), "rotate row"):
		case strings.HasPrefix(scanner.Text(), "rotate column"):
		}
	}
}

func main() {

}
