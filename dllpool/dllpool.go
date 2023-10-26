package dllpool

import (
	"sync"
)

var nodePool = &sync.Pool{
	New: func() interface{} {
		return &node{}
	},
}

type node struct {
	value interface{}
	next  *node
	prev  *node
}

type Dll struct {
	head   *node
	tail   *node
	length int
}

func New() *Dll {
	return &Dll{}
}

func (dll *Dll) Len() int {
	return dll.length
}

func (dll *Dll) PushRight(val interface{}) {
	node := nodePool.Get().(*node)
	node.value = val
	node.next = nil
	node.prev = dll.tail

	if dll.length != 0 {
		dll.tail.next = node
		dll.tail = node
	} else {
		dll.head = node
		dll.tail = node
	}
	dll.length++
}

func (dll *Dll) PushLeft(val interface{}) {
	node := nodePool.Get().(*node)
	node.value = val
	node.next = dll.head
	node.prev = nil

	if dll.length != 0 {
		dll.head.prev = node
		dll.head = node
	} else {
		dll.head = node
		dll.tail = node
	}
	dll.length++
}

func (dll *Dll) PopRight(numPop ...int) interface{} {
	if dll.tail == nil {
		return nil
	}

	num := 1
	if len(numPop) > 0 {
		num = numPop[0]
	}
	if num > dll.length {
		num = dll.length
	}

	if num == 1 {
		temp := dll.tail
		poppedValue := temp.value
		dll.tail = dll.tail.prev
		if dll.tail != nil {
			dll.tail.next = nil
		} else {
			dll.head = nil // list becomes empty after this pop
		}
		dll.length--

		temp.value = nil
		temp.next = nil
		temp.prev = nil
		nodePool.Put(temp)
		return poppedValue
	}

	poppedValues := make([]interface{}, num)
	current := dll.tail
	for i := 0; i < num; i++ {
		poppedValues[i] = current.value // fill the array from most recently popped first (helps in using as stack)
		temp := current
		current = current.prev
		temp.value = nil
		temp.next = nil
		temp.prev = nil
		nodePool.Put(temp)
	}
	dll.tail = current
	if dll.tail != nil {
		dll.tail.next = nil
	} else {
		dll.head = nil // list becomes empty
	}
	dll.length -= num
	return poppedValues
}


func (dll *Dll) PopLeft(numPop ...int) interface{} {
	if dll.head == nil {
		return nil
	}

	num := 1
	if len(numPop) > 0 {
		num = numPop[0]
	}
	if num > dll.length {
		num = dll.length
	}

	if num == 1 {
		// only popping one node
		temp := dll.head
		poppedValue := temp.value
		dll.head = dll.head.next
		temp.next = nil
		temp.value = nil
		if dll.head != nil {
			dll.head.prev = nil
		} else {
			dll.tail = nil
		}
		dll.length--
		return poppedValue
	}

	poppedValues := make([]interface{}, num)
	current := dll.head
	for i := 0; i < num; i++ {
		poppedValues[i] = current.value // in case you want to implement deque as stack
		temp := current
		current = current.next
		temp.value = nil
		temp.prev = nil
		temp.next = nil
		nodePool.Put(temp)
	}
	dll.head = current
	if dll.head != nil {
		dll.head.prev = nil
	} else {
		dll.tail = nil
	}
	dll.length -= num
	return poppedValues
}


func (dll *Dll) PeekHead() interface{} {
	// nil check to avoid nil pointer dereference for empty dll
	if dll.head != nil {
		return dll.head.value
	}
	return nil
}

func (dll *Dll) PeekTail() interface{} {
	if dll.tail != nil {
		return dll.tail.value
	}
	return nil
}