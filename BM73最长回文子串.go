package main

func hasCycleStr(str string) string {
	max := 0
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			v := hasCycleValue(str[i:j])
			if v > max {
				max = v
			}
		}
	}
	return ""
}

func hasCycleValue(num string) int {
	for i := 0; i < len(num)/2; i++ {
		if num[i] != num[len(num)-1-i] {
			return 0
		}
	}
	return len(num)
}

func main() {
	str := "aadddda"
	hasCycleStr(str)
}
