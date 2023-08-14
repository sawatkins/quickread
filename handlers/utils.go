package handlers

import "strconv"

// parse bool from query string
func parseBool(value string) bool {
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false // default to false. is this the best way to handle this?
	}
	return result
}