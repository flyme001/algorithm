给定一个长度为 n 的数组 num 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。

例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。

窗口大于数组长度或窗口长度为0的时候，返回空。

示例1
输入：
[2,3,4,2,6,2,5,1],3
复制
返回值：
[4,4,6,6,6,5]

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
