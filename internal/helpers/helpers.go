package helpers

import "time"

// CheckValueInMap check is the value exist in the given map
func CheckValueInMap(searchIn []string, searchFor string) bool {
	for _, value := range searchIn {
		if searchFor == value {
			return true
		}
	}
	return false
}

// MinDate returns min allowing birthdate
func MinDate() time.Time {
	today := time.Now()
	year, month, day := today.Date()
	startToday := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	startToday = startToday.AddDate(-61, 0, 1)
	return startToday
}

// MaxDate returns max allowing birthdate
func MaxDate() time.Time {
	today := time.Now()
	year, month, day := today.Date()
	startToday := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	startToday = startToday.AddDate(-18, 0, 0)
	return startToday
}
