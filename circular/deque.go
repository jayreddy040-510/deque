package circular

type Deque struct {
    data []interface{}
    front int
    back int
    length int
    capacity int
}

func New(capacity int) *Deque {
    return &Deque{
        data: make([]interface{}, capacity),
        front: 0,
        back: 0,
        length: 0,
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

    pos := (d.back + 1) % d.length
    d.data[pos] = v
    d.length++
}
