package facade

import (
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"
)

type ContainerHandlerFunction func(ctx containerlibcontext.Context, task model.Task) error
