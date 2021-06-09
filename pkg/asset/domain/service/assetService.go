package service

// This component is in charge of handle domain rules
import (
	"context"
	"ms-asset/pkg/asset/domain/entity"
	"ms-asset/pkg/asset/domain/repository"
	"ms-asset/shared/customerror"
)

// assetServiceInterface interface that establishes functions to be implemented
type AssetServiceInterface interface {
	Add(ctx context.Context, b *entity.Asset) error
	GetByClient(ctx context.Context, clientId string) ([]*entity.Asset, error)
	GetBy(ctx context.Context, filters map[string]string) ([]*entity.Asset, error)
	Get(ctx context.Context, sku string) (*entity.Asset, error)
}

type assetService struct {
	repository repository.AssetRepository
}

// New creates a domain service for core logic
func New(repository repository.AssetRepository) AssetServiceInterface {
	return &assetService{repository}
}

// Addasset adds the given record
func (service *assetService) Add(ctx context.Context, p *entity.Asset) error {
	_, error := service.repository.FetchByID(p.RegistrationNumber)
	if error == customerror.ErrRecordNotFound {
		return service.repository.Create(ctx, p)
	}
	return nil
}

// GetByClient searches all records related
func (service *assetService) GetByClient(ctx context.Context, clientId string) ([]*entity.Asset, error) {
	return service.repository.FetchByClient(clientId)
}

// GetBy searches all records related to the filter indicated
func (service *assetService) GetBy(ctx context.Context, filters map[string]string) ([]*entity.Asset, error) {
	return service.repository.FetchBy(filters)
}

// Getasset searches a record
func (service *assetService) Get(ctx context.Context, sku string) (*entity.Asset, error) {
	return service.repository.FetchByID(sku)
}
