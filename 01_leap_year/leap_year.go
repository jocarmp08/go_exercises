package leapyear

// IsLeapYear receives an integer representing a year and
// returns a boolean true if the year is leap, or false otherwise.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
