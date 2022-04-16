package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(years int) bool {
	if years % 4 == 0 {
		if years % 100 == 0{
			if years % 400 == 0{
				return true
			}
			return false
		}
		return true
	}
	return false
}
