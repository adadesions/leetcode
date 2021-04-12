package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (n *Node) add(data int) {
	newNode := Node{data, nil}
	for n != nil {
		if n.next == nil {
			n.next = &newNode
			break
		}
		n = n.next
	}
}

func (n *Node) remove(data int) {
	for n != nil {
		if n.next.data == data {
			if n.next.next == nil {
				n.next = nil
			} else {
				n.next = n.next.next
			}
			break
		}
		n = n.next
	}
}

func (n *Node) update(find int, replace int) {
	for n != nil {
		if n.data == find {
			n.data = replace
		}
		n = n.next
	}
}

func (n *Node) pop() *Node {
	for n != nil {
		if n.next.next == nil {
			n.next = nil
			return n
		}
		n = n.next
	}
	return nil
}

func (n *Node) print() {
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func main() {
	head := Node{5, nil}

	head.add(1)
	head.add(2)
	head.add(3)

	head.remove(2)
	head.update(3, 10)
	head.pop()

	head.print()
}
