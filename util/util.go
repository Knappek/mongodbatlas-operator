package util

import (
	"fmt"
	"strconv"
)

// StringToFloat64 transforms n to float64 and panics if error
func StringToFloat64(s string) float64 {
	if s == "" {
		return 0.0
	}
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}
	return f
}

// Float64ToString transforms a float64 to string
func Float64ToString(f float64) string {
	return fmt.Sprintf("%f", f)
}

// IsNotEqual does not purely compare equality like reflect.DeepEqual does
// It returns false if a and b are equal
// It returns false if a and b are not equal but a is nil
// It returns true if a and b are not equal and a is not nil
func IsNotEqual(a, b interface{}) bool {
	if a != b {
		return !IsZeroValue(a)
	}
	return false
}

// IsZeroValue returns true if input interface is the corresponding zero value
func IsZeroValue(i interface{}) bool {
	if i == nil {
		return true
	} // nil interface
	if i == "" {
		return true
	} // zero value of a string
	if i == 0.0 {
		return true
	} // zero value of a float64
	if i == 0 {
		return true
	} // zero value of an int
	if i == false {
		return true
	} // zero value of a boolean
	return false
}
