package main

func hasCycleString(str string) bool {
	if str != "" {
		return false
	}

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
