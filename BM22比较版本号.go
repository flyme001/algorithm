package main

import (
	"strconv"
	"strings"
)

func compare(version1, version2 string) bool {

	v1 := strings.Trim(version1, ".")
	v1 = strings.Trim(v1, "0")

	v2 := strings.Trim(version2, ".")
	v2 = strings.Trim(v2, "0")
	v1new, _ := strconv.Atoi(v1)
	v2new, _ := strconv.Atoi(v2)

	if v1new > v2new {
		return true
	}
	return false
}

func main() {
	version1 := ""
	version2 := ""
	compare(version1, version2)

}
