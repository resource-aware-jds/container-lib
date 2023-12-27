package model

type Task struct {
	ID         string
	Attributes TaskAttributes
}

type TaskAttributes []byte
