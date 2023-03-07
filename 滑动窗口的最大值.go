package main

func slid(nums []int, index int) []int {
	res := []int{}

	if len(nums) == 0 {
		return []int{}
	}

	for i := 0; i < len(nums); i++ {
		max := getMax(nums[i : index+i])
		res = append(res, max)
	}

	return res
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
