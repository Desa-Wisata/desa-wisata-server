package usecase

import (
	"context"

	"github.com/desa-wisata/library/response"
)

func (uc *usecase) GetDestination(ctx context.Context) response.Message {
	ctx, dest, err := uc.repo.GetDestination(ctx)
	if err != nil {
		return response.InternalError(ctx)
	}

	return response.Success(ctx, dest)
}
