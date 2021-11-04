package easy

func IsPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	for n > 3 {
		if n%3 != 0 {
			return false
		}

		n /= 3
	}

	return n >= 3
}
