package ring

type ring struct {
	data   []interface{}
	head   int
	tail   int
	length int
}

// Ring::Add
func (r *ring) Add(item interface{}) {
	if r.length == cap(r.data) {
		panic("overflow")
	}
	r.data[r.tail] = item
	r.tail++
	if r.tail == cap(r.data) {
		r.tail = 0
	}
	r.length++
}

// Ring::Remove
func (r *ring) Remove() interface{} {
	if r.length == 0 {
		panic("underflow")
	}
	item := r.data[r.head]
	r.data[r.head] = nil
	r.head++
	if r.head == r.length {
		r.head = 0
	}
	r.length--
	if r.length == 0 {
		r.head = 0
		r.tail = 0
	}
	return item
}

// Ring::Peek
func (r *ring) Peek(n int) interface{} {
	if r.length == 0 || n < 0 || n > r.length {
		panic("range error")
	}
	h := r.head + n
	if h > cap(r.data) {
		h -= cap(r.data)
	}
	return r.data[h]
}

// Ring::Clear
func (r *ring) Clear() {
	for ; r.length > 0; r.length-- {
		r.data[r.head] = nil
		r.head++
		if r.head == r.length {
			r.head = 0
		}
	}
	r.head = 0
	r.tail = 0
}

// Ring::Len
func (r *ring) Len() int {
	return r.length
}

// Ring::Cap
func (r *ring) Cap() int {
	return cap(r.data)
}

// Ring::Empty
func (r *ring) Empty() bool {
	return r.length == 0
}

// ring::AtCapacity
func (r *ring) AtCapactiy() bool {
	return r.length == cap(r.data)
}
