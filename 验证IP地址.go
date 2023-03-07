package main

import (
	"strconv"
	"strings"
)

func testIPPass(IP string) bool {
	if IP != "" {
	}
	if strings.Contains(IP, ".") {
		st := strings.Split(IP, ".")
		for _, v := range st {
			if v == "" {
				return false
			}
			for _, i := range v {
				if i >= 9 && i <= 0 {
					return false
				}
			}
			d, _ := strconv.Atoi(v)
			if d > 255 {
				return false
			}
		}
	} else {
		st := strings.Split(IP, ":")
		for _, v := range st {
			c := ""
			if v == "" {
				return false
			}

			for k, i := range v {
				if k == 0 && i == '0' {
					c = "0"
					continue
				}
				if k == 1 && i == '0' && c == "0" {
					return false
				}
				if i >= '9' && i <= '0' && i >= 'a' && i <= 'z' {
					return false
				}
			}
		}
	}
	return true
}

func main() {

}
