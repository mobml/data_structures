//package list implemts a doubly linked list
package list

import (
	"errors"
)

type Node struct {
	info any
	next *Node
	prev *Node
}

type LinkedList struct {
	lenght int
	head   *Node
	tail   *Node
}

func New() *LinkedList {
	return &LinkedList{
		lenght: 0,
		head:   nil,
		tail:   nil,
	}
}

func (node *Node) Value() (any, error) {
	if node == nil {
		err := errors.New("The node is nil")

		return nil, err
	}
	return node.info, nil
}

func newNode(info any) *Node {
	return &Node{
		info: info,
		next: nil,
		prev: nil,
	}
}

func (list *LinkedList) Len() int {
	return list.lenght
}

func (list *LinkedList) PushBack(info any) {
	node := newNode(info)

	if list.head == nil {
		list.head = node
	} else {
		node.prev = list.tail
		list.tail.next = node
	}
	list.tail = node
	list.lenght++
}

func (list *LinkedList) PushFront(info any) {
	node := newNode(info)

	if list.head == nil {
		list.tail = node
	} else {
		list.head.prev = node
		node.next = list.head
	}

	list.head = node
	list.lenght++
}

func (list *LinkedList) Front() *Node {
	if list.lenght == 0 {
		return nil
	}

	return list.head

}

func (list *LinkedList) Back() *Node {
	if list.lenght == 0 {
		return nil
	}
	return list.tail
}

func (node *Node) Next() *Node {
	return node.next
}