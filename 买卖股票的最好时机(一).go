package main

func maxProfit(num []int) int {
	max := make(map[string]int)

	for i := 0; i < len(max); i++ {
		for j := i + 1; j < len(max)-1; j++ {
			max["max"] = num[0]
			if num[j] > num[i] {
				if num[j]-num[i] > max["max"] {
					max["max"] = num[j] - num[i]
				}
			}
		}
	}
	return max["max"]
}

func main() {

}
