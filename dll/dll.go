package dll

type node struct {
	value interface{}
	next  *node
	prev  *node
}

type Deque struct {
	head   *node
	tail   *node
	length int
}

func New() *Deque {
	return &Deque{}
}

func (dll *Deque) Len() int {
	return dll.length
}

func (dll *Deque) PushBack(val interface{}) {
	node := &node{
		value: val,
		next:  nil,
		prev:  dll.tail,
	}

	if dll.length != 0 {
		dll.tail.next = node
		dll.tail = node
	} else {
		dll.head = node
		dll.tail = node
	}
	dll.length++
}

func (dll *Deque) PushFront(val interface{}) {
	node := &node{
		value: val,
		next:  dll.head,
		prev:  nil,
	}

	if dll.length != 0 {
		dll.head.prev = node
		dll.head = node
	} else {
		dll.head = node
		dll.tail = node
	}
	dll.length++
}

func (dll *Deque) PopBack(numPop ...int) interface{} {
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
		poppedValue := dll.tail.value
		dll.tail = dll.tail.prev
		if dll.tail != nil {
			dll.tail.next = nil
		} else {
			dll.head = nil // list becomes empty after this pop
		}
		dll.length--
		return poppedValue
	}

	poppedValues := make([]interface{}, num)
	current := dll.tail
	for i := 0; i < num; i++ {
		poppedValues[i] = current.value // fill the array from most recently popped first (helps in using as stack)
		current = current.prev
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


func (dll *Deque) PopFront(numPop ...int) interface{} {
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
		poppedValue := dll.head.value
		dll.head = dll.head.next
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
		current = current.next
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


func (dll *Deque) PeekHead() interface{} {
	// nil check to avoid nil pointer dereference for empty dll
	if dll.head != nil {
		return dll.head.value
	}
	return nil
}

func (dll *Deque) PeekTail() interface{} {
	if dll.tail != nil {
		return dll.tail.value
	}
	return nil
}
