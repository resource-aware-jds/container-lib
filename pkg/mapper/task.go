package mapper

import (
	"errors"
	"github.com/resource-aware-jds/container-lib/generated/proto/github.com/resource-aware-jds/container-lib/generated/proto"
	"github.com/resource-aware-jds/container-lib/model"
	"strings"
)

var (
	ErrInvalidTaskProto = errors.New("task proto is invalid")
)

func ConvertTaskProtoToModel(taskProto *proto.Task) (*model.Task, error) {
	id := taskProto.GetID()
	attributes := taskProto.GetTaskAttributes()

	idSplit := strings.Split(id, ":")
	if len(idSplit) != 2 {
		return nil, ErrInvalidTaskProto
	}

	return &model.Task{
		ID: model.TaskID{
			JobID:  idSplit[0],
			TaskID: idSplit[1],
		},
		Attributes: attributes,
	}, nil
}
