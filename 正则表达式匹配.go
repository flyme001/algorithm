package main

import "strings"

func match(str1, str2 string) {
	str := []string{}
	for i := 0; i < len(str1); i++ {
		for j := i + 1; j < len(str1)-1; j++ {
			str = append(str, str1[i:j])
		}
	}

	for _, v := range str {
		matchStr(v, str2)
	}
}

func matchStr(str, str2 string) {
	for i := 0; i < len(str2); i++ {
		if str2[i] == str[i]
	}
}

func main() {
	str1 := "aaa"
	str2 := "aba"
	match(str1, str2)
}
