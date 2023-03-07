package main

func noSameArr(num []int) []int {
	v := make(map[int]interface{})
	for _, v1 := range num {
		if _, ok := v[v1]; !ok {
			v[v1] = ""
		}
	}
	data := []int{}
	for k := range v {
		data = append(data, k)
	}
	return data
}
