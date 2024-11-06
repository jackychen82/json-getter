package jsongetter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func GetNodeValue(path string, in []byte) any {
	if path == "" || len(in) == 0 {
		return nil
	}

	input, err := jsonToMap(in)
	if err != nil || input == nil {
		return nil
	}

	paths := strings.Split(path, ".")
	if len(paths) == 1 {
		if v, ok := input[paths[0]]; ok {
			return v
		}
		return nil
	}

	next, ok := input[paths[0]]
	if !ok {
		return nil
	}

	n, err := json.Marshal(next)
	if err != nil {
		return nil
	}

	return GetNodeValue(strings.Join(paths[1:], "."), n)
}

func jsonToMap(in []byte) (map[string]any, error) {
	var m = map[string]any{}
	var out any
	if err := json.Unmarshal(in, &out); err != nil {
		return nil, err
	}

	kind := reflect.TypeOf(out).Kind()

	switch kind {
	case reflect.Array, reflect.Slice:
		outSlice := reflect.ValueOf(out)
		for i := 0; i < outSlice.Len(); i++ {
			m[fmt.Sprintf("[%d]", i)] = outSlice.Index(i).Interface()
		}
		return m, nil
	case reflect.Map:
		if m, ok := out.(map[string]interface{}); ok {
			return m, nil
		}
	}

	return nil, nil
}
