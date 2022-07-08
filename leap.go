// Leap - an exercism exercise
package leap

// IsLeapYear report if it is a leap year.
func IsLeapYear(y int) bool {
	return y%100 != 0 && y%4 == 0 || y%400 == 0
}
