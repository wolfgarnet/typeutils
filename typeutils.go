package typeutils

import (
	"reflect"
	"github.com/wolfgarnet/logging"
)

var logger logging.Logger

func FindField(instance interface{}, field string, depth int) interface{} {
	logger.Debug("Instance: %v, Field: %v, depth: %v", instance, field, depth)

	value := reflect.Indirect(reflect.ValueOf(instance))

	for i := 0 ; i < value.NumField() ; i++ {
		fieldValue := value.Field(i)
		fieldType := value.Type().Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			if depth > 0 || depth == -1 {
				result := FindField(fieldValue.Interface(), field, (depth-1))
				if result != nil {
					return result
				}
			}

		default:
			if fieldType.Name == field {
				return fieldValue.Interface()
			}
		}


	}

	return nil
}