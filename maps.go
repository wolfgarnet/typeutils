package typeutils

import "reflect"

func MergeMaps(src, dst map[string]interface{}, overwrite bool) map[string]interface{} {
	for k, v := range src {
		_, ok := dst[k]

		value := reflect.Indirect(reflect.ValueOf(v))
		switch value.Kind() {
		case reflect.Map:

		}

		if !ok || overwrite {
			dst[k] = v
			continue
		}

	}

	return dst
}