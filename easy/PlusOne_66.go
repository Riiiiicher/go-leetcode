package easy

func PlusOne(digits []int) []int {
	remainder := 0
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + 1
		if tmp > 9 {
			remainder = tmp / 10
			digits[i] = tmp % 10
		} else {
			digits[i] = tmp
			return digits
		}
	}
	if remainder > 0 {
		digits = append([]int{remainder}, digits...)
	}
	return digits
}
