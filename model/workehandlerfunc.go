package model

import "github.com/resource-aware-jds/worker-lib/pkg/workerlibcontext"

type WorkerHandlerFunc func(ctx workerlibcontext.Context) error
