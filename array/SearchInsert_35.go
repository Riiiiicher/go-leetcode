package array

func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if target == nums[m] {
			return m
		}
		if target > nums[m] {
			if m == r || target < nums[m+1] {
				return m + 1
			}
			l = m + 1
		} else {
			if m == 0 || target > nums[m-1] {
				return m
			}
			r = m - 1
		}
	}
	return -1
}
