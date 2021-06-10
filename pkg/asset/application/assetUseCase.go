package application

import (
	"context"
	"ms-asset/pkg/asset/domain/entity"
	"ms-asset/pkg/asset/domain/service"
	"ms-asset/pkg/asset/infrastructure/mapper"
	"ms-asset/pkg/asset/infrastructure/request"
)

// assetUseCaseInterface provides operations to be executed.
type AssetUseCaseInterface interface {
	Add(ctx context.Context, requestData request.AssetRequest) error
	GetByClient(ctx context.Context, clientId string) ([]*entity.Asset, error)
	GetBy(ctx context.Context, filters map[string]string) ([]*entity.Asset, error)
	Get(ctx context.Context, sku string) (*entity.Asset, error)
}

type assetUseCase struct {
	service service.AssetServiceInterface
}

// New creates the application use case class from App Layer
func New(service service.AssetServiceInterface) AssetUseCaseInterface {
	return &assetUseCase{service}
}

// Add adds the given record to storage
func (app *assetUseCase) Add(ctx context.Context, requestData request.AssetRequest) error {
	asset := entity.New(itemsRequestToDomain(requestData.Images), requestData.Furnished, requestData.RentingPrice, requestData.Area, requestData.Rooms,
		requestData.BathRooms, requestData.Parkings, requestData.Country, requestData.Province, requestData.City,
		requestData.Description, requestData.Category, requestData.State, requestData.Type, requestData.RegistrationNumber)
	return app.service.Add(ctx, asset)

}

// GetByClient searches all records into the storage
func (app *assetUseCase) GetByClient(ctx context.Context, clientId string) ([]*entity.Asset, error) {
	return app.service.GetByClient(ctx, clientId)
}

// GetBy searches all records into the storage
func (app *assetUseCase) GetBy(ctx context.Context, filters map[string]string) ([]*entity.Asset, error) {
	filtersMapped, _ := mapper.New().RequestToFilter(filters)
	return app.service.GetBy(ctx, filtersMapped)
}

// Get searches all records into the storage
func (app *assetUseCase) Get(ctx context.Context, sku string) (*entity.Asset, error) {
	return app.service.Get(ctx, sku)
}

func itemsRequestToDomain(images []request.AssetImageRequest) []entity.AssetImage {
	var results []entity.AssetImage
	for _, image := range images {
		results = append(results, entity.AssetImage{Url: image.Url,
			Description: image.Description, Type: image.Type, State: image.State})
	}
	return results
}
