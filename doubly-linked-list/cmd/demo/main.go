package main

import (
	"github.com/NutcrackerCom/go-backend-journey/doubly-linked-list/dll"
)

func main() {
	var list dll.List
	list.Push_back("one")
	list.Push_back("two")
	list.Push_back("three")
	list.Push_back("four")
	list.PrintNode()
}
