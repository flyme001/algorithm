package main

func maxArr(arr []int) []int {
	l := len(arr)
	max := 0
	m := make(map[string][]int, 0)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l-1; j++ {
			s := sum(arr[i:j])
			if s > max {
				max = s
				m["max"] = arr[i:j]
			}
		}
	}

	return m["max"]
}

func sum(num []int) int {
	s := 0
	for i := 0; i < len(num); i++ {
		s = s + num[0]
	}
	return s
}

func main() {

}
