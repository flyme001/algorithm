package main

func isvalid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	m := make(map[byte]byte, 0)
	m['('] = ')'
	m['{'] = '}'
	m['['] = ']'

	t := []byte(s)
	for k, v := range []byte(s) {
		if k == 0 {
			continue
		}

		if m[t[len(t)-1-k]] == v {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {

}
