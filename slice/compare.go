package utils

import (
	"reflect"
)

// EqualsUnordered compares two slices without failing if they're not in the same order
func EqualsUnordered(x []interface{}, y []interface{}) bool {
	if x == nil && y != nil {
		return false
	}
	if x != nil && y == nil {
		return false
	}

	reflect.DeepEqual()

	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)

	if v1.Type() != v2.Type() {
		return false
	}
	if len(x) != len(y) {
		return false
	}
	switch v1.Type() {

	}

	}
}
