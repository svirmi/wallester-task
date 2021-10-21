package enums

type Gender int

const (
	Male Gender = iota
	Female
)

var names = [...]string{"Male", "Female"}

func (gender Gender) String() string {
	if gender < Male || gender > Female {
		return "Unknown"
	}
	return names[gender]
}

func Exists(requiredName string) bool {
	for _, name := range names {
		if requiredName == name {
			return true
		}
	}
	return false
}
