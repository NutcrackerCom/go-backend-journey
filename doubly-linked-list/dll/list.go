package dll

import (
	"fmt"
)

/*

prev    curr    next
  |      |        |
[nil]  [node]  [node]

*/

type Node struct {
	next *Node
	prev *Node
	val  string
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (list *List) Push_back(str string) {
	var node Node = Node{
		next: nil,
		prev: nil,
		val:  str,
	}
	if list.head == nil {
		list.head = &node
	} else {
		list.tail.next = &node
		node.prev = list.tail
	}
	list.tail = &node
	list.size++
}

func (list *List) Push_front(str string) {
	var node Node = Node{
		next: nil,
		prev: nil,
		val:  str,
	}
	if list.head == nil {
		list.tail = &node
	} else {
		list.head.prev = &node
		node.next = list.head
	}
	list.head = &node
	list.size++
}

func (list *List) PrintNode() {
	current := list.head
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func (list *List) Pop_front() {
	if list.head != nil {
		list.head = list.head.next
		list.size--
	}
}

func (list *List) Pop_back() {
	if list.size == 1 {
		list.head = nil
		list.tail = nil
		list.size--
	}
	if list.head != nil {
		list.tail = list.tail.prev
		list.tail.next = nil
		list.size--
	}
}

func (list *List) Insert(position int, str string) {
	if position > list.size || position < 0 {
		return
	}
	switch position {
	case 0:
		list.Push_front(str)
	case list.size:
		list.Push_back(str)
	default:
		var node Node = Node{
			next: nil,
			prev: nil,
			val:  str,
		}
		current := list.head
		for i := 0; i < position-1; i++ {
			current = current.next
		}
		next := current.next
		current.next = &node
		node.prev = current
		node.next = next
		next.prev = &node
		list.size++
	}
}

func (list *List) Erase(position int) {
	if position >= list.size || position < 0 {
		return
	}
	switch position {
	case 0:
		list.Pop_front()
	case list.size - 1:
		list.Pop_back()
	default:
		current := list.head
		for i := 0; i < position; i++ {
			current = current.next
		}
		next := current.next
		current.prev.next = next
		next.prev = current.prev
		list.size--
	}
}

func (list *List) Revert() {
	if list.size <= 1 {
		return
	}
	newTail := list.head
	current := list.head
	var prev *Node = nil
	var next *Node = nil
	for current != nil {
		next = current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	list.head = prev
	list.tail = newTail
}

func (list *List) Copy() List {
	var copyList List
	if list.size == 0 {
		return copyList
	}
	current := list.head
	for current != nil {
		copyList.Push_back(current.val)
		current = current.next
	}
	return copyList
}

func (list *List) GetSize() int {
	return list.size
}
