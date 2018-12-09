package main

import "fmt"

var Debug = false

func main() {
	Debug = true
	fmt.Println("32", Do2(9, 25))
	Debug = false
	fmt.Println("8317", Do2(10, 1618))
	fmt.Println("146373", Do2(13, 7999))
	fmt.Println("2764", Do2(17, 1104))
	fmt.Println("54718", Do2(21, 6111))
	fmt.Println("37305", Do2(30, 5807))

	fmt.Println(Do(458, 72019))
	fmt.Println(Do2(458, 72019))
	fmt.Println(Do2(458, 72019*100))

}

func Do(players, max int) int {
	scores := make([]int, players)
	currentPlayer := 0
	index := 0
	r := []int{0}

	for i := 1; i < max; i++ {
		if i%23 == 0 {
			index -= 7
			if index < 0 {
				index += len(r)
			}
			scores[currentPlayer] += i
			scores[currentPlayer] += r[index]
			r = append(r[:index], r[index+1:]...)
		} else {
			index++
			if index == len(r) {
				index = 0
			}
			r = append(r[:index+1], append([]int{i}, r[index+1:]...)...)
			index++
		}
		currentPlayer++
		if currentPlayer == len(scores) {
			currentPlayer = 0
		}
		if i%72019 == 0 {
			fmt.Println("%")
		}
	}
	top := 0
	for _, p := range scores {
		if p > top {
			top = p
		}
	}
	return top
}

func Do2(players, max int) int {
	scores := make([]int, players)
	currentPlayer := 0
	root := &Node{Value: 0}
	root.Next = root
	root.Prev = root
	index := root

	for i := 1; i <= max; i++ {
		if i%23 == 0 {
			for i := 0; i < 7; i++ {
				index = index.Prev
				if Debug {
					fmt.Println(index.Value)
				}
			}
			scores[currentPlayer] += i
			scores[currentPlayer] += index.Value
			index = index.Remove()
		} else {
			index = index.Next
			index = index.Append(i)
		}
		currentPlayer++
		if currentPlayer == len(scores) {
			currentPlayer = 0
		}
		if Debug {
			root.Print(index)
		}
	}
	top := 0
	for _, p := range scores {
		if p > top {
			top = p
		}
	}
	return top
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func (n *Node) Append(val int) *Node {
	nx := &Node{Value: val}
	nx.Next = n.Next
	nx.Prev = n
	n.Next.Prev = nx
	n.Next = nx
	return nx
}
func (n *Node) Remove() *Node {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	return n.Next
}

func (n *Node) Print(index *Node) {
	fmt.Print("[", n.Value, ",")
	current := n.Next
	for current != n {
		if current == index {
			fmt.Print("(", current.Value, "),")
		} else {
			fmt.Print(current.Value, ",")
		}
		current = current.Next
	}
	fmt.Println("]")
}
