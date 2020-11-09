package urlparams

import (
	"regexp"
	"strings"
)

// GetParams receives a string representing an url and
// returns a map containing the parameters from the query
// string and its corresponding values
//
// Assume the url syntax is always correct
func GetParams(url string) map[string]string {
	params := make(map[string]string)

	var re = regexp.MustCompile(`(?m)[^&?]*?=[^&?]*`)
	for _, match := range re.FindAllString(url, -1) {
		values := strings.Split(match, "=")
		params[values[0]] = values[1]
	}

	return params
}
