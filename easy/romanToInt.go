package easy

func RomanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0
	lastInt := 0
	for i := len(s) - 1; i >= 0; i = i - 1 {
		num := romanMap[s[i]]
		if num < lastInt {
			ans = ans - num
		} else {
			ans = ans + num
		}
		lastInt = num
	}
	return ans
}
