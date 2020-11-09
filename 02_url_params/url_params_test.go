package urlparams

import (
	"reflect"
	"testing"
)

func TestGetParamsForRegularCase(t *testing.T) {
	url := "http://localhost/app/mod?param1=val1&param2=val2&param3=val3"
	expectedResult := map[string]string{
		"param1": "val1",
		"param2": "val2",
		"param3": "val3",
	}

	result := GetParams(url)

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Error: Expected result %v, Actual result %v", expectedResult, result)
	}
}

func TestGetParamsForNoParamsCase(t *testing.T) {
	url := "http://localhost/app/mod"
	expectedResult := make(map[string]string)
	result := GetParams(url)

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Error: Expected result %v, Actual result %v", expectedResult, result)
	}
}
