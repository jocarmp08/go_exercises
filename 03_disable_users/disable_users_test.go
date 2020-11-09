package disableusers

import (
	"reflect"
	"testing"

	utilities "../internal"
)

func TestDisableUsersRegularCase(t *testing.T) {
	payload := `[
		{"usr":"diego.maradona","enabled":true,"expires":"13-05-2022"}, 
		{"usr":"neymarjr","enabled":true,"expires":"05-12-2024"}, 
		{"usr":"bryan.ruiz","enabled":false,"expires":"23-12-2018"}, 
		{"usr":"lsuarez","enabled":true,"expires":"03-07-2019"}
	]`

	expectedResult, _ := utilities.JSONStringToArrayOfMaps(`[
		{"usr":"diego.maradona","enabled":true,"expires":"13-05-2022"},
		{"usr":"neymarjr","enabled":true,"expires":"05-12-2024"},
		{"usr":"bryan.ruiz","enabled":false,"expires":"23-12-2018"},
		{"usr":"lsuarez","enabled":false,"expires":"03-07-2019"}
	]`)

	result, _ := utilities.JSONStringToArrayOfMaps(DisableUsers(payload))

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Error: Expected result %v, Actual result %v", expectedResult, result)
	}
}

func TestDisableUsersEmptyList(t *testing.T) {
	payload := `[]`

	expectedResult, _ := utilities.JSONStringToArrayOfMaps(`[]`)

	result, _ := utilities.JSONStringToArrayOfMaps(DisableUsers(payload))

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Error: Expected result %v, Actual result %v", expectedResult, result)
	}
}
