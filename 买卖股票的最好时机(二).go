package main

func maxMoeny(num []int) int {
	sum := 0
	for i := 0; i < len(num)-1; {
		for j := i + 1; j < len(num)-1; j++ {
			if num[i] < num[j] {
				sum = sum + num[j] - num[i]
				i = j
			}
		}
	}
	return sum
}

func main() {
}
