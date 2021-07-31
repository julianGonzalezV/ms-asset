package application

import (
	"context"
	"ms-asset/pkg/asset/domain/service"
	"ms-asset/pkg/asset/infrastructure/mapper"
	rest "ms-asset/pkg/asset/infrastructure/request"
	"ms-asset/pkg/asset/infrastructure/response"
)

// assetUseCaseInterface provides operations to be executed.
type AssetUseCaseInterface interface {
	Add(ctx context.Context, requestData rest.AssetRequest) error
	GetByClient(ctx context.Context, clientId string) ([]response.AssetFullResponse, error)
	GetBy(ctx context.Context, filters map[string]string) ([]response.AssetFullResponse, error)
	Get(ctx context.Context, code string) (response.AssetFullResponse, error)
}

type assetUseCase struct {
	service service.AssetServiceInterface
}

// New creates the application use case class from App Layer
func New(service service.AssetServiceInterface) AssetUseCaseInterface {
	return &assetUseCase{service}
}

// Add adds the given record to storage
func (app *assetUseCase) Add(ctx context.Context, requestData rest.AssetRequest) error {
	return app.service.Add(ctx, mapper.New().RequestToDomain(requestData))
}

// Get ByClient searches all records into the storage
func (app *assetUseCase) GetByClient(ctx context.Context, clientId string) ([]response.AssetFullResponse, error) {
	assetL, error := app.service.GetByClient(ctx, clientId)
	if error != nil {
		return nil, error
	}
	return mapper.New().AssetDomainToFullResponseList(assetL), nil
}

// GetBy searches all records into the storage
func (app *assetUseCase) GetBy(ctx context.Context, filters map[string]string) ([]response.AssetFullResponse, error) {
	filtersMapped, _ := mapper.New().RequestToFilter(filters)
	assetL, error := app.service.GetBy(ctx, filtersMapped)
	if error != nil {
		return nil, error
	}
	return mapper.New().AssetDomainToFullResponseList(assetL), nil
}

// Get searches all records into the storage
func (app *assetUseCase) Get(ctx context.Context, code string) (response.AssetFullResponse, error) {
	asset, error := app.service.Get(ctx, code)
	if error == nil {
		return mapper.New().DomainToFullResponse(*asset), nil
	}
	return response.AssetFullResponse{}, error
}
