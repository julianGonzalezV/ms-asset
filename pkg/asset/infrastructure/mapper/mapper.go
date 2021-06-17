package mapper

import (
	"strconv"
	"strings"
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
