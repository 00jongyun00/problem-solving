package easy

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {
		if closer, ok := open[r]; ok {
			stack = append(stack, closer)
			continue
		}

		l := len(stack) - 1
		if l < 0 || r != stack[l] {
			return false
		}
		stack = stack[:l]
	}
	return len(stack) == 0
}
