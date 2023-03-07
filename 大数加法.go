package main

import "strconv"

func sumString(num1, num2 string) string {
	if num1 == "" {
		return num2
	}
	if num2 == "" {
		return num1
	}
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)

	return strconv.Itoa(n1 + n2)

}
