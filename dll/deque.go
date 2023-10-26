package dll

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

func (dll *Dll) PushLeft(val interface{}) {
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

func (dll *Dll) PopRight(numPop ...int) interface{} {
	if len(numPop) == 0 {
		if dll.length != 0 {
			temp := dll.tail
			dll.tail = dll.tail.prev
			if dll.tail != nil {
				dll.tail.next = nil
			} else {
				dll.head = nil // list becomes empty after this pop
			}
			dll.length--
			temp.prev = nil
			return temp.value
		}
		return nil
	} else {
		num := numPop[0]
		poppedValues := make([]interface{}, 0, num)
		if num > dll.length {
			num = dll.length
		}
		for i := 0; i < num; i++ {
			if dll.tail != nil {
				poppedValues = append(poppedValues, dll.tail.value)
				prev := dll.tail.prev
				dll.tail.prev = nil
				dll.tail = prev
				if dll.tail != nil {
					dll.tail.next = nil
				} else {
					dll.head = nil // list becomes empty
				}
				dll.length--
			}
		}
		return poppedValues
	}
}

func (dll *Dll) PopLeft(numPop ...int) interface{} {
	head := dll.head
	if len(numPop) == 0 {
		if head != nil {
			dll.head = head.next
			dll.length--
			return head.value
		} else {
			return nil
		}
	} else {
		num := numPop[0]
		if num > dll.length {
			num = dll.length
		}
		poppedValues := make([]interface{}, 0, num)
		for i := 0; i < num; i++ {
			if head != nil {
				poppedValues = append(poppedValues, head.value)
				head = head.next
				head.prev = nil	
			} else {
				dll.tail = nil
			}
			dll.length--
		}
		return poppedValues
	}
}
