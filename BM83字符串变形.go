package main

import "strings"

func changeST(str string) string {

	st := ""
	t := []byte{}
	if strings.Contains(str, " ") {
		for i := 0; i < len(str); i++ {
			st = st + string(str[len(str)-i])
		}
		t = []byte(st)
	} else {
		t = []byte(st)
	}

	st1 := ""
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
