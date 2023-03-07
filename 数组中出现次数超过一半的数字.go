package main

func moreNums(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	half := len(nums) / 2
	val := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := val[nums[i]]; !ok {
			val[nums[i]] = 0
		} else {
			val[nums[i]] = val[nums[i]] + 1
		}
	}

	for _, v := range val {
		if v >= half {
			return v
		}
		continue
	}

	return -1
}

func main() {

}
