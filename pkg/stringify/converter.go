package stringify

import "strconv"

// Helper functions to convert types to strings
func BoolToString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

func ToInteger(num int) string {
	return strconv.Itoa(num)
}
