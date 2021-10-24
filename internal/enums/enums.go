package enums

type Gender int

const (
	Male Gender = iota
	Female
)

var names = [...]string{"Male", "Female"}

// String overrides string method for the enum and returns value depends of given key
func (gender Gender) String() string {
	if gender < Male || gender > Female {
		return "Unknown"
	}
	return names[gender]
}

// Exists checks is the given value exists in the enum
func Exists(requiredName string) bool {
	for _, name := range names {
		if requiredName == name {
			return true
		}
	}
	return false
}
