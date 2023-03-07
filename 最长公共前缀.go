package main

import "strings"

func longStringSame(num []string) string {
	min := ""
	minLen := 0
	for _, v := range num {
		if len(v) > minLen {
			min = v
		}
	}

	diff := false
	c := 0
	for m := 0; m < len(min); m++ {
		tmp := strings.Split(min, "")[m]

		for _, v := range num {
			if strings.Split(v, "")[m] == tmp {
				c = c + 1
			} else {
				diff = true
			}
		}
		if diff {
			break
		}
	}

	return min[:c]
}
