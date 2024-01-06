package mapper

import (
	"errors"
	"github.com/resource-aware-jds/container-lib/generated/proto/github.com/resource-aware-jds/container-lib/generated/proto"
	"github.com/resource-aware-jds/container-lib/model"
)

var (
	ErrInvalidTaskProto = errors.New("task proto is invalid")
)

func ConvertTaskProtoToModel(taskProto *proto.Task) (*model.Task, error) {

	return &model.Task{
		ID:    taskProto.GetID(),
		JobID: taskProto.GetJobID(),

		Attributes: taskProto.GetTaskAttributes(),
	}, nil
}
