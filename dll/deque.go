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
			dll.length--
			return temp.value
		} else {
			return nil
		}
	} else {
		num := numPop[0]
		if num > dll.length {
			num = dll.length
		}

	}
}