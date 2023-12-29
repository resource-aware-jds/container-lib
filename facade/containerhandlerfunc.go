package facade

import "github.com/resource-aware-jds/worker-lib/pkg/containerlibcontext"

type ContainerHandlerFunction func(ctx containerlibcontext.Context) error
