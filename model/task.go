package model

type Task struct {
	ID         TaskID
	Attributes TaskAttributes
}

type TaskAttributes []byte

type TaskID struct {
	JobID  string
	TaskID string
}
