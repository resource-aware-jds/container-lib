package model

type Task struct {
	ID         string
	JobID      string
	Attributes TaskAttributes
}

type TaskAttributes []byte
