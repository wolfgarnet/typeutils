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

// Traverse embedded structs given a predicate: check.
// Requiring each struct implements Super() to return its super type/struct.
func Traverse(obj interface{}, check func(reflect.Type) bool) (reflect.Type) {
	tp := reflect.TypeOf(obj)
	logger.Debug("tp: %v", tp.Name())

	for {
		if check(tp) {
			return tp
		}

		// Try super
		m, ok := tp.MethodByName("Super")

		if !ok {
			break;
		}

		in := make([]reflect.Value,0)
		val2 := m.Func.Call(in)
		tp = val2[0].Type()
	}

	return nil
}

// TypeImplements can determine if an object implements a given method name.
// Requiring each struct implements Super() to return its super type/struct.
func TypeImplements(obj interface{}, name string) (t reflect.Type) {
	tp := reflect.TypeOf(obj)
	logger.Debug("tp: %v", tp.Name())
	t = nil

	for {

		_, ok := tp.MethodByName(name)

		if !ok {
			// Try super
			sm, sok := tp.MethodByName("Super")

			if !sok {
				break;
			}


			in := make([]reflect.Value,0)
			//val2 := sm.Call(in)
			val2 := sm.Func.Call(in)
			tp = val2[0].Type()
		} else {
			println(1)
			t = tp
			return
		}

		break
	}

	return
}