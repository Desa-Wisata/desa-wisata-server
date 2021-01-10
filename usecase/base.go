package usecase

import (
	"context"

	"github.com/desa-wisata/desa-wisata-server/repository"
	"github.com/desa-wisata/library/response"
)

// Base ...
type Base interface {
	GetDestination(ctx context.Context) response.Message
}

type usecase struct {
	repo repository.Base
}

// InitUsecase ...
func InitUsecase(r repository.Base) Base {
	return &usecase{r}
}
