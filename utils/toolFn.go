package utils

import "strings"

func SameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

// camel to snake lower case
func CamelToSnake(s string) string {
	var ret string
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' {
			ret += "_"
		}
		ret += strings.ToLower(string(s[i]))
	}
	return ret
}
