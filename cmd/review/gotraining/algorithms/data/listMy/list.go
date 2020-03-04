package listMy

import "errors"
// todo finish some function in the future
type node struct{
	data string
	next *node
	prev *node
}

type List struct{
	count int
	first *node
	last *node
}

func(l *List)Add(data string)(*node){
	n := node{
		data:data,
	}

	l.count++

	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	l.last.next = &n

	n.prev = l.last

	l.last = &n

	return &n
}

func(l *List)AddFront(data string)(*node){
	n := node{
		data:data,
	}

	l.count++

	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	l.first.prev = &n

	n.next = l.first

	l.first = &n

	return &n
}

func(l *List)Find(data string)(*node,error){

	n := l.first

	for {
		if n.data == data{
			return n,nil
		}
		if n.next == nil {
			return nil,errors.New("not match")
		}
		n = n.next
	}
}

func(l *List)FindResver(data string)(*node,error){
	n := l.last

	for {
		if n.data == data {
			return n,nil
		}
		if n.prev == nil {
			return nil,errors.New("not match")
		}
		n = n.prev
	}
}

func(l *List)Remove(data string)(*node,error){
	if l.last == nil || l.first == nil{
		return nil,errors.New("list is empty")
	}

	n := l.last
	for{
		if n.data == data {
			n.prev.next = n.next
			if n.next != nil{
				n.next.prev = n.prev
			}
			n.prev = nil
			n.next = nil
			return n,nil
		}
		n = n.prev
	}
}

