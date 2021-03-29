package main

import "fmt"

type Node interface {
	SetValue(v int)
	GetValue() int
}

//type SLLNode
type SLLNode struct {
	next  *SLLNode
	value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

//type PowerNode
type PowerNode struct {
	next  *PowerNode
	value int
}

func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

//

//
type level int

func test() {
	sl := new(level)
	sl.raisseShieldLevel(4)
	sl.raisseShieldLevel(5)
}

func (lv *level) raisseShieldLevel(i int) {
	if *lv == 0 {
		*lv = 1
	}
	*lv = (*lv) * level(i)
}

//

//
//type linked list
type SingleLinkedList struct {
	head *SLLNode
	tail *SLLNode
}

func newSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}
func (list *SingleLinkedList) Add(v int) {
	newNode := &SLLNode{value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
}

func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.head; n != nil; n = n.next {
		s += fmt.Sprintf(" {%d} ", n.GetValue())
	}
	return s
}

// Section 3 > Video 11
// methods and interfaces{}
func main() {
	var node Node
	node = NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of value ", node.GetValue())

	node = NewPowerNode()
	node.SetValue(5)
	fmt.Println("Node is of value ", node.GetValue())

	if n, ok := node.(*PowerNode); ok {
		fmt.Println(n.value)
	}

	test()

	// for the example Section 3 > Code Files > singlelikedlist.go
	list := newSingleLinkedList()
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	fmt.Println("Hello, playground", list)
}
