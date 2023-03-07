package main

func findPeakElement(num []int) {

	index := []int{}
	tmp := make(map[int]int)

	for k, v := range num {
		tmp[k] = v
		if v < tmp[k-1] {
			index = append(index, k)
		}
	}
}

func main() {
}
