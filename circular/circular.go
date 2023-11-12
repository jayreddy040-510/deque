package circular

import (
	"fmt"
	"math"
	"strings"
)

type DequeConfig struct {
	initialCapacity int
	minCapacity     int
	shrinkThreshold float64 // ratio of deque length:capacity that triggers dealloc of excess array mem
	shrinkFactor    float64 // factor by which underlying array shrinks when length:capacity <= shrinkThreshold
	growThreshold   float64 // ratio of deque length:capacity that triggers underlying array growing logic
	growFactor      float64 // factor by which deque grows when length:capacity >= growThreshold
}

type validationError struct {
	errors []string
}

func (v *validationError) addError(s string) {
	v.errors = append(v.errors, s)
}

func (v *validationError) hasErrors() bool {
	return len(v.errors) > 0
}

func (v *validationError) createErrorMsg() string {
	var b strings.Builder
	b.WriteString("your deque config returned the following validation errors:\n")
	for _, err := range v.errors {
		b.WriteString(fmt.Sprintf("\t%v\n", err))
	}
	return b.String()
}

func (dc *DequeConfig) validate() error {
	ve := &validationError{}
	initialCap := dc.initialCapacity
	minCap := dc.minCapacity
	shrinkThreshold := dc.shrinkThreshold
	shrinkFactor := dc.shrinkFactor
	growThreshold := dc.growThreshold
	growFactor := dc.growFactor

	// initial capacity > 0
	if initialCap < 1 {
		ve.addError(fmt.Sprintf(
			"initial capacity: %d must be greater than 0",
			initialCap,
		))
	}

	// min capacity > 0
	if minCap < 1 {
		ve.addError(fmt.Sprintf(
			"initial capacity: %d must be greater than 0",
			minCap,
		))
	}

	// initial capacity < min capacity
	if initialCap < minCap {
		ve.addError(fmt.Sprintf(
			"initial capacity: %d cannot be smaller than min capacity: %d",
			initialCap,
			minCap,
		))
	}

	// 0 < growThreshold < 1
	if growThreshold <= 0 || growThreshold >= 1 {
		ve.addError(fmt.Sprintf(
			"grow threshold: %v must be between 0 and 1 exclusive",
			growThreshold,
		))
	}

	// 0 < shrinkThreshold < 1
	if shrinkThreshold <= 0 || shrinkThreshold >= 1 {
		ve.addError(fmt.Sprintf(
			"shrink threshold: %v must be between 0 and 1 exclusive",
			shrinkThreshold,
		))
	}

	// growThreshold > shrinkThreshold
	if growThreshold <= shrinkThreshold {
		ve.addError(fmt.Sprintf(
			"grow threshold: %v must be greater than shrink threshold: %v",
			growThreshold,
			shrinkThreshold,
		))
	}

	// growFactor > 1
	if growFactor <= 1 {
		ve.addError(fmt.Sprintf(
			"growFactor: %v must be greater than 1",
			growFactor,
		))
	}

	// shrinkFactor < 1
	if shrinkFactor >= 1 {
		ve.addError(fmt.Sprintf(
			"shrinkFactor: %v must be lesser than 1",
			shrinkFactor,
		))
	}

	// growFactor > growThreshold
	if growFactor <= growThreshold {
		ve.addError(fmt.Sprintf(
			"growFactor: %v must be greater than growThreshold: %v",
			growFactor,
			growThreshold,
		))
	}

	// shrinkFactor < shrinkThreshold
	if shrinkFactor >= shrinkThreshold {
		ve.addError(fmt.Sprintf(
			"shrinkFactor: %v must be lesser than shrinkThreshold: %v",
			shrinkFactor,
			shrinkThreshold,
		))
	}

	if ve.hasErrors() {
		return fmt.Errorf("%v", ve.createErrorMsg())
	} else {
		return nil
	}
}

type Deque struct {
	data     []interface{}
	front    int
	back     int
	length   int
	capacity int
	config   *DequeConfig
}

var defaultConfig = &DequeConfig{
	initialCapacity: 10,
	minCapacity:     10,
	shrinkThreshold: 0.25,
	shrinkFactor:    0.5,
	growThreshold:   1.0,
	growFactor:      2.0,
}

func New(config ...*DequeConfig) (*Deque, error) {
	if len(config) == 0 {
		return &Deque{
			data:     make([]interface{}, defaultConfig.initialCapacity),
			front:    0,
			back:     0,
			length:   0,
			capacity: defaultConfig.initialCapacity,
			config:   defaultConfig,
		}, nil
	} else {
		if err := config[0].validate(); err != nil {
			return nil, err
		}
		return &Deque{
			data:     make([]interface{}, config[0].initialCapacity),
			front:    0,
			back:     0,
			length:   0,
			capacity: config[0].initialCapacity,
			config:   config[0],
		}, nil
	}
}

func (d *Deque) resize(size ...int) {
	var newCapacity int

	if len(size) == 0 {
		newCapacity = int(math.Round(d.config.growFactor * float64(d.capacity)))
	} else {
		newCapacity = size[0]
	}

	newData := make([]interface{}, newCapacity)
	pos := 0

	for i := 0; i < d.length; i++ {
		newData[pos] = d.data[(d.front+i)%d.capacity]
		pos++
	}

	d.front = 0
	d.back = pos
	d.data = newData
	d.capacity = newCapacity
}

func (d *Deque) PushBackOne(v interface{}) {
	if d.length >= int(math.Round(d.config.growFactor * float64(d.capacity))) {
		d.resize()
	}

	if d.length != 0 {
		d.back = (d.back + 1) % d.capacity
	}

	d.data[d.back] = v
	d.length++
}

func (d *Deque) PushBackBulk(values []interface{}) {
	if finalLength := d.length + len(values); finalLength > int(math.Round(d.config.growThreshold * float64(d.capacity))) {
		// im ashamed, lol - is there a better way to multiply a float into an int,
		// preserve the data by rounding and then result in an int? need int for capacity in resize()
		d.resize(int(math.Round(d.config.growFactor * float64(finalLength)))) 
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

	for i := 0; i < n; i++ {
		popped[i] = d.data[d.back]
		d.data[d.back] = nil
		d.back = (d.back - 1 + d.capacity) % d.capacity
	}
	d.length -= n
	if d.length == 0 {
		d.front, d.back = 0, 0
	}

	// inline shrinking logic: if ratio of length:capacity < shrinkThreshold set in config struct &&
	// if deque capacity > minimum capacity, resize to either min capacity, or shrinkFactor * deque capacity
	if float64(d.length)/float64(d.capacity) <= d.config.shrinkThreshold && d.capacity > d.config.minCapacity {
		if shrunkCapacity := int(math.Round(d.config.shrinkFactor * float64(d.capacity))); shrunkCapacity > d.config.minCapacity {
			d.resize(shrunkCapacity)
		} else {
			d.resize(d.config.minCapacity)
		}
	}
	return popped
}

func (d *Deque) PushFrontOne(v interface{}) {
	if d.length >= int(math.Round(d.config.growFactor * float64(d.capacity))) {
		d.resize()
	}

	if d.length != 0 {
		d.front = (d.front - 1 + d.capacity) % d.capacity
	} else {
		d.front, d.back = 0, 0
	}

	d.data[d.front] = v
	d.length++
}

func (d *Deque) PushFrontBulk(values []interface{}) {
	if finalLength := len(values) + d.length; finalLength > int(math.Round(d.config.growThreshold * float64(d.capacity))) {
		d.resize(int(math.Round(d.config.growFactor * float64(finalLength))))
	}

	for i := 0; i < len(values); i++ {
		d.front = (d.front - 1 + d.capacity) % d.capacity
		d.data[d.front] = values[i]
	}

	d.length += len(values)
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

	if float64(d.length)/float64(d.capacity) <= d.config.shrinkThreshold && d.capacity > d.config.minCapacity {
		if shrunkCapacity := int(math.Round(d.config.shrinkFactor * float64(d.capacity))); shrunkCapacity > d.config.minCapacity {
			d.resize(shrunkCapacity)
		} else {
			d.resize(d.config.minCapacity)
		}
	}

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
