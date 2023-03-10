package main

func reverseStr(str string) string {
	data := []rune(str)
	for i := 0; i < len(data)/2; i++ {
		data[i] = data[len(str)-1-i]
		data[len(str)-1-i] = data[i]
	}
	return string(data)
}
