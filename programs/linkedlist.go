package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (l *Node) Insert(val int) {
	newNode := &Node{
		Value: val,
		Next:  nil,
	}

	for l.Next != nil {
		l = l.Next
	}

	l.Next = newNode
}

func (l *Node) DeleteFromEnd() {

	if l.Next == nil {
		l = nil
		return
	}

	for l.Next.Next != nil {
		l = l.Next
	}

	l.Next = nil
}

func (l *Node) PrintList() {

	for l != nil {
		fmt.Println(l.Value)
		l = l.Next
	}

}

func main() {

	Head := Node{
		Value: 7,
		Next:  nil,
	}

	// Head.Insert(5)
	// Head.Insert(6)
	// Head.Insert(10)
	// Head.PrintList()
	// fmt.Println(Head)
	Head.DeleteFromEnd()
	// fmt.Println()
	fmt.Println(Head)
	Head.PrintList()
}
