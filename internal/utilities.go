package utilities

import "encoding/json"

// JSONStringToArrayOfMaps ...
func JSONStringToArrayOfMaps(jsonString string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// IsElementInSlice ...
func IsElementInSlice(element int, slice []int) bool {
	return _binarySearch(slice, 0, len(slice)-1, element) > -1
}

// BinarySearch ...
func _binarySearch(slice []int, l int, r int, element int) int {
	if l > r {
		return -1
	}

	m := (l + r) / 2

	if slice[m] == element {
		return m
	} else if slice[m] > element {
		return _binarySearch(slice, l, m-1, element)
	} else {
		return _binarySearch(slice, m+1, r, element)
	}
}
