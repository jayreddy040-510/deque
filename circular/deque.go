package circular

type Deque struct {
    data []interface{}
    front int
    back int
    size int
}

func New(capacity int) *Deque {
    return &Deque{
        data: make([]interface{}, capacity),
        front: 0,
        back: 0,
        size: capacity,
    }
}

func (d *Deque) resize() {
    newSize := 2 * d.size 

    newData := make([]interface{}, newSize)


}

func (d *Deque) PushRightOne(v interface{}) {
    if len(d.data) >= d.size {
        d.resize()
    }
}
