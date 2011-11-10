package queue

// queue.Interface
type Interface interface {
	Add       (interface{})
	Remove    ()    interface{}
	Peek      (int) interface{}
	Clear     ()
	Len       () int
	Empty     () bool
	Cap       () int
	AtCapactiy() bool
}

func NewQueue(capacity int) Interface {
	return &queue{data: make([]interface{}, capacity), head: 0, tail: 0, length: 0}
}

