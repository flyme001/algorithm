package main

import "strings"

func changeST(str string) string {
	st := ""
	for i := 0; i < len(str); i++ {
		st = st + string(str[len(str)-i])
	}

	st1 := ""
	t := []byte(st)

	for _, v := range t {
		if v >= 'a' && v <= 'z' {
			st1 = st1 + strings.ToUpper(string(v))
		} else {
			st1 = st1 + strings.ToLower(string(v))
		}
	}

	return st1
}

func main() {

}
