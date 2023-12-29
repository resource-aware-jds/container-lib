package facade

import "github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"

type ContainerHandlerFunction func(ctx containerlibcontext.Context) error
