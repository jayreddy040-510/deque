package dll               

import (
	"fmt"
)

type node struct {
	value interface{}
	next *node
	prev *node
}

type Dll struct {
	head *node
	tail *node
	length int
}

func New() *Dll {
	return &Dll{}
}

func (dll *Dll) Len() int{
	return dll.length
}

func (dll *Dll) PushRight(val interface{}) {
	node := &node{
		value: val,
		next: nil,
		prev: dll.tail,
	}

	if dll.tail != nil {
		dll.tail.next = node
	}

	if dll.length == 0 {
		dll.head = node
	}

	dll.tail = node
	dll.length++
}

func (dll *Dll) PushLeft(val interface{}) {
	node := &node{
		value: val,
		next: dll.head,
		prev: nil,
	}

	if dll.head != nil {
		dll.head.prev = node
	}	

	if dll.length == 0 {
		dll.tail = node
	}

	dll.head = node
	dll.length++
}