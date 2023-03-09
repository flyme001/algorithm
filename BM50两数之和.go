package main

func twoSum(nums []int, target int) []int {
	res := []int{}
	v := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		v[nums[i]] = i
	}

	for j := 0; j < len(nums); j++ {
		if _, ok := v[target-nums[j]]; ok {
			res = append(res, j, v[target-nums[j]])
			return res
		}
	}

	return []int{}
}

func main() {
	nums := []int{3, 5, 6}
	twoSum(nums, 9)
}
