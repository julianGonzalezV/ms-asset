package mapper

import (
	"strconv"
)

type Mapper interface {
	RequestToFilter(queryParameters map[string]string) (map[string]interface{}, error)
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
		if value != "" {
			v, err := typeConverter(value)
			if err == nil {
				result[key] = v
			}
		}
	}
	return result, nil
}

func typeConverter(v string) (interface{}, error) {
	f, error := strconv.ParseFloat(v, 64)
	if error != nil {
		return nil, error
	}
	return f, nil

}
