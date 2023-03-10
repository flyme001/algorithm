package main

import "strconv"

func sumString(num1, num2 string) string {
	if num1 == "" {
		return num2
	}
	if num2 == "" {
		return num1
	}

	flag := false
	for i := 0; i < len(num1); i++ {
		n1 := int(num1[len(num1)-1-i])
		if len(num2)-1-i > 0 {
			n2 := int(num1[len(num2)-1-i])
			if n1+n2 > 10 {
				flag = true
			} else {
				flag = false
			}
		}

	}

	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)

	return strconv.Itoa(n1 + n2)

}

func main() {

	strA := "aaaa"
	strB := "dd"

	if len(strA) > len(strB) {
		sumString(strA, strB)
	} else {
		sumString(strB, strA)
	}

}
