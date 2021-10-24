package helpers

// CheckValueInMap check is the value exist in the given map
func CheckValueInMap(searchIn []string, searchFor string) bool {
	for _, value := range searchIn {
		if searchFor == value {
			return true
		}
	}
	return false
}
