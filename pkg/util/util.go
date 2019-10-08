package util

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