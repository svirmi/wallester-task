package helpers

func CheckValueInMap(searchIn []string, searchFor string) bool {
	for _, value := range searchIn {
		if searchFor == value {
			return true
		}
	}
	return false
}
