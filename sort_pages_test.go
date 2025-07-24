package main

import (
	"reflect"
	"testing"
)

func TestGetSortedKeys(t *testing.T) {
	tests := []struct {
		name     string
		inputMap map[string]int
		expected []page
	}{
		{
			name: "simple sort",
			inputMap: map[string]int{
				"e": 4,
				"b": 2,
				"a": 1,
				"c": 3,
				"f": 1,
				"d": 3,
			},
			expected: []page{
				{"e", 4},
				{"c", 3},
				{"d", 3},
				{"b", 2},
				{"a", 1},
				{"f", 1},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.inputMap)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected slice: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
