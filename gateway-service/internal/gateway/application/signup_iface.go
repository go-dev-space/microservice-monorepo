package application

import (
	"context"

	"github.com/microservice-monorepo/gateway-service/internal/gateway/interfaces/dto"
)

type Registrar interface {
	Execute(ctx context.Context, p *dto.Payload) (map[string]string, error)
}
