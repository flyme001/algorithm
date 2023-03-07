package main

func allocateLeader(num [][]int) {
	data := []int{}
	for _, v := range num {
		for _, v1 := range v {
			data = append(data, v1)
		}
	}

	sum := 1
	for i := 1; i < len(data)-2; i++ {
		if data[i] > data[i+1] {
			sum = sum + 1
		}
	}

}
