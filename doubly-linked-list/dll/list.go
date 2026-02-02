package dll

import "fmt"

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

func (list *List) push_back(str string) {
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

func (list *List) push_front(str string) {
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

func (list *List) pop_front() {
	if list.head != nil {
		list.head = list.head.next
		list.size--
	}
}

func (list *List) pop_back() {
	if list.head != nil {
		list.tail = list.tail.prev
		list.tail.next = nil
		list.size--
	}
}

func (list *List) insert(position int, str string) {
	if position > list.size {
		return
	}
	if position == 0 {
		list.push_front(str)
	} else if position == list.size {
		list.push_back(str)
	} else {
		var node Node = Node{
			next: nil,
			prev: nil,
			val:  str,
		}
		current := list.head
		for i := 0; i < position; i++ {
			current = current.next
		}
		next := current.next
		current.next = &node
		node.prev = current
		node.next = next
		next.prev = &node
	}
}
