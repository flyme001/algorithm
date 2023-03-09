package main

func slid(nums []int, index int) int {
	max := 0

	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums)-3; i++ {
		val := getMax(nums[i : index+i])
		if max < val {
			max = val
		}
	}

	return max
}

func getMax(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	nums := []int{1, 4, 2, 5, 7, 8}
	slid(nums, 3)
}
