package main

func insert(arr []int) []int {
	res := []int{}

	if len(arr) == 0 {
		return []int{}
	}

	if len(arr) == 1 {
		res = append(res, arr[0])
		return res
	}

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			continue
		}

		if (i+1)%2 == 0 {
			l := (i+1)/2 - 1
			r := (i + 1) / 2
			res = append(res, GetMedian(arr[l:r]))
		} else {
			res = append(res, arr[i+1/2])
		}
	}

	return []int{}
}

func GetMedian(data []int) int {
	return int(data[0] + data[1]/2)
}

func main() {

}
