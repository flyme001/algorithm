给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

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
