package main
//https://www.youtube.com/watch?v=sYcOK51hl-A

type Node struct{
	next *Node
	value interface{}
}

func Reverse(Header *Node)(){
	var current *Node
	var prev *Node
	var next *Node

	current = Header.next
	prev = nil
	for current!=nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	Header = prev
}
