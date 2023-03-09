package main

func firstNums(nums []int) (res []int) {
	if len(nums) == 0 {
		return []int{}
	}

	val := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := val[nums[i]]; !ok {
			val[nums[i]] = 0
		} else {
			val[nums[i]] = val[nums[i]] + 1
		}
	}

	for _, v := range val {
		if v == 1 {
			res = append(res, v)
		}
		continue
	}

	return res
}

func main() {

}

