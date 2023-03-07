package main

func hasCylcle(str string) string {
	l := 0
	
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str)-1; j++ {
			v := hasCycleValue(str[i:j])
			if v > l {
				l = v
			}
		}
	}

	return ""
}

func hasCycleValue(num string) int {
	for i := 0; i < len(num)/2; i++ {
		if num[i] != num[len(num)-i] {
			return 0
		}
	}
	return len(num)
}

func main() {
	str := "adddddda"
	hasCylcle(str)
}
