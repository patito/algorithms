package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	last  *Node
	head  *Node
	count int
}

func New() *LinkedList {
	list := &LinkedList{}
	return list
}

func (l *LinkedList) IsEmpty() bool {
	return l.count == 0
}

func (l *LinkedList) Length() int {
	return l.count
}

func (l *LinkedList) Preppend(data interface{}) {
	node := &Node{
		data: data,
	}
	if l.IsEmpty() {
		l.last = node
	}
	node.next = l.head
	l.head = node
	l.count++
}

func (l *LinkedList) Append(data interface{}) {
	node := &Node{
		data: data,
	}
	if l.IsEmpty() {
		l.head = node
		l.last = node
	} else {
		l.last.next = node
		l.last = node
	}
	l.count++
}

func (l *LinkedList) Insert(data interface{}, pos int) {
	if pos <= 0 {
		l.Preppend(data)
	} else if pos >= l.Length() {
		l.Append(data)
	} else {
		current := l.head
		previous := l.head
		counter := 0
		node := &Node{
			data: data,
		}
		for current != nil {
			if pos == counter {
				previous.next = node
				node.next = current
				l.count++
				return
			}
			previous = current
			current = current.next
			counter++
		}
	}
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (l *LinkedList) Search(data interface{}) bool {
	current := l.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

func (l *LinkedList) Remove(data interface{}) {

	if l.head.data == data {
		l.head = l.head.next
		l.count--
	} else {
		current := l.head
		previous := l.head
		for current != nil {
			if current.data == data {
				previous.next = current.next
				l.count--
				return
			}
			previous = current
			current = current.next
		}
	}
}

func (l *LinkedList) Pop(pos int) {
	if l.IsEmpty() {
		return
	}

	if pos <= 0 {
		l.head = l.head.next
		l.count--
	} else {
		current := l.head
		previous := l.head
		counter := 0
		for current != nil {
			if counter == pos {
				previous.next = current.next
				l.count--
				return
			}
			previous = current
			current = current.next
			counter++
		}
	}
}

/*
func main() {
	list := New()

	list.Insert("end", 3)
	list.Insert("init", -1)
	list.Insert("middle", 1)
	list.Print()
	fmt.Println(list.Length())

}*/
