package main

func robbery(arr []int) int {
	suma := 0
	sumb := 0
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			suma = suma + arr[i]
		} else {
			sumb = sumb + arr[i]
		}
	}

	if suma > sumb {
		return suma
	}

	return sumb
}
