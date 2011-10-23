package queue

type Queue interface {
	Add       (interface{})
	Remove    ()    interface{}
	Peek      (int) interface{}
	Clear     ()
	Len       () int
	Empty     () bool
	Cap       () int
	AtCapactiy() bool
}

func NewQueue(capacity int) Queue {
	return &queue{data: make([]interface{}, capacity), head: 0, tail: 0, length: 0}
}

