package main

import "strings"

func sameString(num1, num2 string) string {
	if num1 == "" || num2 == "" {
		return ""
	}

	str := []string{}
	for i := 0; i < len(num1); i++ {
		for j := i + 1; j < len(num1)-1; j++ {
			temp := num1[i:j]
			str = append(str, temp)
		}
	}

	max := 0
	strMax := ""
	for x := 0; x < len(str); x++ {
		if strings.Contains(num2, str[x]) && len(str[x]) > max {
			max = len(str[x])
			strMax = str[x]
		}
	}

	return strMax
}

func main() {
	num1 := "1AB2345CD"
	num2 := "12345EF"
	sameString(num1, num2)
}
