package repository

import (
	"context"
	"ms-asset/pkg/asset/domain/entity"
)

// AssetRepository interface that establishes functions to be implemented
type AssetRepository interface {
	// Create saves a new record
	Create(ctx context.Context, p *entity.Asset) error
	// FetchByID returns the record with given ID
	FetchByID(Id string) (*entity.Asset, error)
	// FetchByClient returns the record with given an ID
	FetchByClient(clientId string) ([]*entity.Asset, error)
	// FetchByClient returns the record with given any filter in filters map
	FetchBy(filters map[string]string) ([]*entity.Asset, error)
}
