package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
)

// EvaluateContext contains context information for Resource to handle CRUD
type EvaluateContext struct {
	WriteService services.WriteService
}

// Intent contains Intent Compiler state for a resource.
type Intent interface {
	Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error
}

// BaseIntent implements the default Evaluate interface
type BaseIntent struct {
}

// Evaluate creates/updates/deletes lower-level resources when needed.
func (b *BaseIntent) Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error {
	return nil
}
