package typeutils

import (
	"testing"
	"reflect"
)

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		src, dst, expected map[string]interface{}
		overwrite bool
	}{
		{
			map[string]interface{}{"a":"b"},
			map[string]interface{}{"a":"b"},
			map[string]interface{}{"a":"b"},
			false,
		},
		{
			map[string]interface{}{"a":"b"},
			map[string]interface{}{"a":"c"},
			map[string]interface{}{"a":"c"},
			false,
		},
		{
			map[string]interface{}{"a":"b"},
			map[string]interface{}{"c":"d"},
			map[string]interface{}{"a":"b", "c": "d"},
			false,
		},
		{
			map[string]interface{}{"a":"b"},
			map[string]interface{}{"a":"c"},
			map[string]interface{}{"a":"b"},
			true,
		},
		{
			map[string]interface{}{"a":map[string]interface{}{"b": "c"}},
			map[string]interface{}{"a":map[string]interface{}{"d": "e"}},
			map[string]interface{}{"a":map[string]interface{}{"b": "c", "d": "e"}},
			false,
		},
	}

	for i, test := range tests {
		MergeMaps(test.src, test.dst, test.overwrite)

		//logger.Debug("RES %v: %v", i, test.dst)

		if !reflect.DeepEqual(test.dst, test.expected) {
			t.Errorf("Maps #%v does not match:\n%v\n%v", i, test.dst, test.expected)
		}

	}
}