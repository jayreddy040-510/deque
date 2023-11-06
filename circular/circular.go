package circular

type Deque struct {
	data     []interface{}
	front    int
	back     int
	length   int
	capacity int
}

func New(capacity int) *Deque {
	return &Deque{
		data:     make([]interface{}, capacity),
		front:    0,
		back:     0,
		length:   0,
		capacity: capacity,
	}
}

func (d *Deque) resize(size ...int) {
	var newCapacity int

	if len(size) == 0 {
		newCapacity = 2 * d.capacity
	} else {
		newCapacity = size[0]
	}

	newData := make([]interface{}, newCapacity)
	pos := 0

	for i := 0; i < d.length; i++ {
		newData[pos] = d.data[d.front]
		d.front = (d.front + 1) % d.capacity
		pos++
	}

	d.front = 0
	d.back = pos
	d.data = newData
	d.capacity = newCapacity
}

func (d *Deque) PushBackOne(v interface{}) {
	if d.length >= d.capacity {
		d.resize()
	}

	if d.length != 0 {
		d.back = (d.back + 1) % d.capacity
	}

	d.data[d.back] = v
	d.length++
}

func (d *Deque) PushBackBulk(values []interface{}) {
	if finalLength := d.length + len(values); finalLength > d.capacity {
		d.resize(2 * finalLength)
	}

	for _, v := range values {
		if d.length != 0 {
			d.back = (d.back + 1) % d.capacity
		}

		d.data[d.back] = v
	}
	d.length += len(values)
}

func (d *Deque) PopBackOne() interface{} {
	if d.length == 0 {
		return nil
	}

	popped := d.data[d.back]
	d.data[d.back] = nil
	if d.length != 1 {
		d.back = (d.back - 1 + d.capacity) % d.capacity
	} else {
		d.back, d.front = 0, 0
	}
	d.length--
	return popped
}

func (d *Deque) PopBackBulk(n int) []interface{} {
	if n > d.length {
		n = d.length
	}

	popped := make([]interface{}, n)

	for i :=0; i < n; i++ {
		popped[i] = d.data[d.back]
		d.data[d.back] = nil
		d.back = (d.back - 1 + d.capacity) % d.capacity
	}
	d.length -= n
	if d.length == 0 {
		d.front, d.back = 0, 0
	}
	return popped
}

func (d *Deque) PushFrontOne(v interface{}) {
	if d.length >= d.capacity {
		d.resize()
	}

	if d.length != 0 {
		if d.front != 0 {
			d.front--
		} else {
			d.front = d.capacity - 1
		}
	} else {
		d.front, d.back = 0, 0
	}

	d.data[d.front] = v
	d.length++
}

func (d *Deque) PopFrontOne() interface{} {
	if d.length == 0 {
		return nil
	}

	popped := d.data[d.front]
	d.data[d.front] = nil

	if d.length != 1 {
		d.front = (d.front + 1) % d.capacity
	} else {
		d.front, d.back = 0, 0
	}

	d.length--
	return popped
}

func (d *Deque) PeekHead() interface{} {
	if d.length == 0 {
		return nil
	} else {
		return d.data[d.front]
	}
}

func (d *Deque) PeekTail() interface{} {
	if d.length == 0 {
		return nil
	} else {
		return d.data[d.back]
	}
}
