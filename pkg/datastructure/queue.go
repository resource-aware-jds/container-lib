package datastructure

type Queue[Data any] struct {
	data []Data
}

func ProvideQueue[Data any](size int) Queue[Data] {
	return Queue[Data]{
		data: make([]Data, 0, size),
	}
}

func (q *Queue[Data]) Pop() (*Data, bool) {
	if q.Empty() {
		return nil, false
	}

	result := q.data[0]
	q.data = q.data[1:]
	return &result, true
}

func (q *Queue[Data]) Push(d Data) {
	q.data = append(q.data, d)
}

func (q *Queue[Data]) Empty() bool {
	return len(q.data) == 0
}
