package jsongetter

import (
	"reflect"
	"testing"
)

func TestGetNodeValue(t *testing.T) {
	type testCase struct {
		name         string
		path         string
		input        []byte
		expected     interface{}
		expectedType reflect.Kind
	}

	tests := []testCase{
		{
			name:     "Simple Path",
			path:     "name",
			input:    []byte(`{"name": "John"}`),
			expected: "John",
		},
		{
			name:     "Nested Path",
			path:     "info.address.city",
			input:    []byte(`{"info": {"address": {"city": "New York"}}}`),
			expected: "New York",
		},
		{
			name:     "Array Path",
			path:     "users.[0].name",
			input:    []byte(`{"users": [{"name": "Alice"}, {"name": "Bob"}]}`),
			expected: "Alice",
		},
		{
			name:     "Array Path 2",
			path:     "[1].name",
			input:    []byte(`[{"name": "Alice"}, {"name": "Bob"}, {"name": "Charlie"}]`),
			expected: "Bob",
		},
		{
			name:     "Boolean Path",
			path:     "info.boolean",
			input:    []byte(`{"info": {"boolean": false}}`),
			expected: false,
		},
		{
			name:     "Number Path",
			path:     "info.number",
			input:    []byte(`{"info": {"number": 2}}`),
			expected: float64(2),
		},
		{
			name:         "Invalid Path",
			path:         "info.address.city",
			input:        []byte(`{"info": {"address": "not an object"}}`),
			expected:     nil,
			expectedType: reflect.Invalid,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := GetNodeValue(tc.path, tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected result %v, but got %v", tc.expected, result)
			}

			if tc.expectedType != reflect.Invalid {
				resultType := reflect.TypeOf(result).Kind()
				if resultType != tc.expectedType {
					t.Errorf("Expected result type %v, but got %v", tc.expectedType, resultType)
				}
			}
		})
	}
}
