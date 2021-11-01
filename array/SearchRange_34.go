package array

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	rangeLeft, rangeRight := -1, -1
	for l <= r {
		m := (r-l)/2 + l
		if nums[m] == target {
			rangeLeft, rangeRight = m, m
			for nums[rangeLeft] == target {
				rangeLeft = rangeLeft - 1
			}
			for nums[rangeRight] == target {
				rangeRight = rangeRight + 1
			}
			rangeLeft = rangeLeft + 1
			rangeRight = rangeRight - 1
			break
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return []int{rangeLeft, rangeRight}
}
