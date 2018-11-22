package slice

import (
	"reflect"
)

// EqualsUnordered compares two slices without failing if they're not in the same order
func EqualsUnordered(x interface{}, y interface{}) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	vx := reflect.ValueOf(x)
	vy := reflect.ValueOf(y)

	// If it isn't an array or a slice, do the usual DeepEqual()
	if vx.Kind() != reflect.Slice && vx.Kind() != reflect.Array {
		return reflect.DeepEqual(x, y)
	}

	// Check if elements are the same type
	if vx.Type().Elem() != vy.Type().Elem() {
		return false
	}

	// Arrays / Slices are not the same if their length differs
	if vx.Len() != vy.Len() {
		return false
	}

	// From this point on, type and element type of array are the same.

	if vx.Len() == 0 {
		return true
	}

	xarr := make([]interface{}, vx.Len())
	yarr := make([]interface{}, vy.Len())
	for i := 0; i < vx.Len(); i++ {
		xarr[i] = vx.Index(i).Interface()
		yarr[i] = vy.Index(i).Interface()
	}

	for _, xel := range xarr {
		contains := false
		for _, yel := range yarr {
			if reflect.DeepEqual(xel, yel) {
				contains = true
			}
		}
		if !contains {
			return false
		}
	}
	return true
}
