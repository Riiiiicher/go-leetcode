package easy

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}

	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true

}
