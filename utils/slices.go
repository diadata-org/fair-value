package utils

// ContainsAny returns true if @s contains @a.
func ContainsAny(s []any, a any) bool {
	for _, item := range s {
		if item == a {
			return true
		}
	}
	return false
}

// ContainsAnySlice returns true if @s1 contains @s2.
func ContainsAnySlice(s1, s2 []any) bool {
	for _, item := range s2 {
		if !ContainsAny(s1, item) {
			return false
		}
	}
	return true
}
