package repository

import (
	"context"

	"github.com/desa-wisata/desa-wisata-server/domain"

	"github.com/jinzhu/gorm"
)

// Base ...
type Base interface {
	GetDestination(ctx context.Context) (context.Context, []domain.Destination, error)
}

type repo struct {
	db *gorm.DB
}

//InitRepository ...
func InitRepository(db *gorm.DB) Base {
	return &repo{db}
}
