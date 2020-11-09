package disableusers

import (
	"encoding/json"
	"time"
)

var customDateLayout = "02-01-2006"

type user struct {
	Usr     string `json:"usr"`
	Enabled bool   `json:"enabled"`
	Expires string `json:"expires"`
}

// DisableUsers receives a JSON string containing a list of objects representing users
// and disable all users that, given the current date, are expired.
//
// Safely assume that the received parameter is always a valid json value.
//
// Returns a JSON string containing all the users received, with the users that are expired marked as disabled.
//
// Note: It could be done with the JSONStringToArrayOfMaps function in utilities package
// but this seemed like a good exercise to practice with Go Structs
func DisableUsers(users string) string {
	var usersData []user

	json.Unmarshal([]byte(users), &usersData)

	if len(usersData) > 0 {
		_disableUsers(usersData)
		e, _ := json.Marshal(usersData)
		return string(e)
	}

	return `[]`
}

func _disableUsers(users []user) {
	currentTime := time.Now()
	for i, user := range users {
		t, _ := time.Parse(customDateLayout, user.Expires)
		if t.Before(currentTime) {
			user.setEnabled(false)
			users[i] = user
		}
	}
}

func (u *user) setEnabled(value bool) {
	u.Enabled = value
}
