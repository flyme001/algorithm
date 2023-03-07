package main

import (
	"fmt"
	"strings"
)

func longStr(str1, str2 string) string {
	str := ""
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Println(string(r[i]))
		if strings.Count(str1, string(r[i])) > 0 {
			str = str + string(r[i])
		}
	}
	fmt.Println(str)
	return str
}

func main() {
	str1 := "1A2C3D4B56"
	str2 := "B1D23A456A"
	longStr(str1, str2)
}
