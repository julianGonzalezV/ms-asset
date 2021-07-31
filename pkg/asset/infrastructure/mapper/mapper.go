package mapper

import (
	"ms-asset/pkg/asset/domain/entity"
	"ms-asset/pkg/asset/infrastructure/request"
	"ms-asset/pkg/asset/infrastructure/response"
	"strconv"
	"strings"
)

type Mapper interface {
	RequestToFilter(queryParameters map[string]string) (map[string]interface{}, error)
	RequestToDomain(requestData request.AssetRequest) *entity.Asset
	DomainToFullResponse(assetEntity entity.Asset) response.AssetFullResponse
	DomainToBasicResponse(assetEntity entity.Asset) response.AssetFullResponse
	AssetDomainToFullResponseList(assets []*entity.Asset) []response.AssetFullResponse
}

type mapperImpl struct {
}

// New creates a mapper instance
func New() Mapper {
	return &mapperImpl{}
}

func (mapper *mapperImpl) RequestToFilter(queryParameters map[string]string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for key, value := range queryParameters {
		if value != "" && value != "0" {
			if strings.Contains(value, ":") {
				splResult := strings.Split(value, ":")
				result[key] = map[string]interface{}{"$gte": typeConverter(splResult[0]), "$lte": typeConverter(splResult[1])}
			} else {
				result[key] = typeConverter(value)
			}

		}
	}
	return result, nil
}

///
func typeConverter(v string) interface{} {
	f, error := strconv.ParseFloat(v, 64)
	if error == nil {
		return f
	}
	b, errb := strconv.ParseBool(v)
	if errb == nil {
		return b
	}
	return v

}

func (mapper *mapperImpl) RequestToDomain(requestData request.AssetRequest) *entity.Asset {
	return entity.New(itemsRequestToDomain(requestData.Images), requestData.Furnished, requestData.VisitorParking, requestData.Elevator,
		requestData.CommunalArea, requestData.Gym, requestData.FloorLevel,
		requestData.RentingPrice, requestData.Area, requestData.Rooms,
		requestData.BathRooms, requestData.Parkings, requestData.Country, requestData.Province, requestData.City,
		requestData.Description, requestData.Category, requestData.State, requestData.Type, requestData.Code, requestData.RegistrationNumber)

}

func (mapper *mapperImpl) DomainToFullResponse(assetEntity entity.Asset) response.AssetFullResponse {
	return response.NewFullResponse(itemsDomainToResponse(assetEntity.Images), assetEntity.Furnished, assetEntity.VisitorParking, assetEntity.Elevator,
		assetEntity.CommunalArea, assetEntity.Gym, assetEntity.FloorLevel,
		assetEntity.RentingPrice, assetEntity.Area, assetEntity.Rooms,
		assetEntity.BathRooms, assetEntity.Parkings, assetEntity.Country, assetEntity.Province, assetEntity.City,
		assetEntity.Description, assetEntity.Category, assetEntity.State, assetEntity.Type, assetEntity.Code, assetEntity.RegistrationNumber)

}

func (mapper *mapperImpl) DomainToBasicResponse(assetEntity entity.Asset) response.AssetFullResponse {
	return response.NewFullResponse(itemsDomainToResponse(assetEntity.Images), assetEntity.Furnished, assetEntity.VisitorParking, assetEntity.Elevator,
		assetEntity.CommunalArea, assetEntity.Gym, assetEntity.FloorLevel,
		assetEntity.RentingPrice, assetEntity.Area, assetEntity.Rooms,
		assetEntity.BathRooms, assetEntity.Parkings, assetEntity.Country, assetEntity.Province, assetEntity.City,
		assetEntity.Description, assetEntity.Category, assetEntity.State, assetEntity.Type, assetEntity.Code, assetEntity.RegistrationNumber)

}

func (mapper *mapperImpl) AssetDomainToFullResponseList(assets []*entity.Asset) []response.AssetFullResponse {
	var results []response.AssetFullResponse
	for _, assetItem := range assets {
		results = append(results, mapper.DomainToFullResponse(*assetItem))
	}
	return results
}

func itemsRequestToDomain(images []request.AssetImageRequest) []entity.AssetImage {
	var results []entity.AssetImage
	for _, image := range images {
		results = append(results, entity.AssetImage{Url: image.Url,
			Description: image.Description, Type: image.Type, State: image.State})
	}
	return results
}

func itemsDomainToResponse(images []entity.AssetImage) []response.AssetImageResponse {
	var results []response.AssetImageResponse
	for _, image := range images {
		results = append(results, response.AssetImageResponse{Url: image.Url,
			Description: image.Description, Type: image.Type, State: image.State})
	}
	return results
}
