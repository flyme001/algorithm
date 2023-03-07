package main

func search(num []int, traget int) int {
	if len(num) == 0 {
		return -1
	}

	low := 0
	high := len(num) - 1
	for low <= high {
		middle := (low + high) / 2
		val := num[middle]
		if val == traget {
			return val
		}
		if traget > val {
			low = middle + 1
		}
		if traget < val {
			high = middle - 1
		}
	}

	return traget

}
func main() {
	ar := []int{1, 3, 5, 8, 8}
	search(ar, 5)
}
