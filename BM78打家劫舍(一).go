package main

func rob(num []int) int {
	i := 0
	sum := 0

	if len(num) == 0 {
		return 0
	}

	if len(num) == 1 {
		return num[0]
	}

	if len(num) == 2 {
		if num[0] > num[1] {
			return num[0]
		} else {
			return num[1]
		}
	}

	if len(num) >= 3 {
		if num[0]+num[2] > num[1] {
			i = 2
		} else {
			i = 1
		}
	}

	for j := i; j < len(num); j = j + 2 {
		sum = sum + num[j]
	}

	return sum
}

func main() {

}
