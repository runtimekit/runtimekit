package backend

import (
	"context"

	"github.com/runtimekit/runtimekit/pkg/api/types"
)

type Backend interface {
	GetDevices(ctx context.Context) []types.Device
}

type backend struct {
}

func NewBackend() Backend {
	return &backend{}
}
