package main

import (
	"fmt"
)

type node[T comparable] struct {
	Data T
	next *node[T]
}

type list[T comparable] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		next: nil,
	}

	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.next == nil {
		l.start.next = &n
		return
	}

	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
}

func (l *list[T]) search(data T) bool {
	if l.start == nil {
		return false
	}

	cur := l.start
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *list[T]) delete(data T) {
	if l.start == nil {
		return
	}
	if l.start.Data == data {
		l.start = l.start.next
		return
	}
	cur := l.start
	for cur.next != nil {
		if cur.next.Data == data {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (l *list[T]) PrintMe() {
	if l.start == nil {
		return
	}
	// print all elements
	for current := l.start; current != nil; current = current.next {
		fmt.Println(current.Data)
	}
}

func main() {
	var myList list[int]
	fmt.Println(myList)
	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)
	myList.search(9)
	myList.delete(3)

	myList.PrintMe()

	// Print all elements
	for {
		fmt.Println("*", myList.start)
		if myList.start == nil {
			break
		}
		myList.start = myList.start.next
	}
}
