package model

import "fmt"

type Task struct {
	ID         TaskID
	Attributes TaskAttributes
}

type TaskAttributes []byte

type TaskID struct {
	JobID  string
	TaskID string
}

func (t *TaskID) GetRawTaskID() string {
	return fmt.Sprintf("%s:%s", t.JobID, t.TaskID)
}
