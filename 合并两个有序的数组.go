package main

func mergeArr(num1, num2 []int) []int {
	m := len(num1)
	n := len(num2)

	i := 0
	j := 0

	res := []int{}
	for i < m && j < n {
		if num1[i] > num2[j] {
			res = append(res, num2[j])
			i++
		} else {
			res = append(res, num2[i])
			j++
		}
	}
	return res

}
