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

func (d *Deque) resize() {
	newCapacity := 2 * d.capacity

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

func (d *Deque) PopBackOne() interface{} {
	if d.length == 0 {
		return nil
	}

	popped := d.data[d.back]
	d.data[d.back] = nil
	if d.length != 1 {
		if d.back != 0 {
			d.back--
		} else {
			d.back = d.capacity - 1
		}
	} else {
		d.back, d.front = 0, 0
	}
	d.length--
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