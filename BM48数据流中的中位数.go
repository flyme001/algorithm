如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。
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
