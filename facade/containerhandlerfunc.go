package facade

import "github.com/resource-aware-jds/worker-lib/pkg/workerlibcontext"

type ContainerHandlerFunction func(ctx workerlibcontext.Context) error
