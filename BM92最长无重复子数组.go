package main

func noSameArr(num []int) int {
	max := 0

	if len(num) == 0 {
		return 0
	}
	if len(num) == 1 {
		return num[0]
	}

	for i := 0; i < len(num); i++ {
		for j := i; j < len(num); j++ {
			if num[i] != num[j] {
				l := len(num[i:j])
				if l > max {
					max = l
				}
			} else {
				break
			}
		}
	}

	return 0
}
