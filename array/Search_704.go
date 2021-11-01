package array

func Search(nums []int, target int) int {
	lIdx, rIdx := 0, len(nums)-1
	for lIdx <= rIdx {
		h := (rIdx-lIdx)/2 + lIdx
		if nums[h] == target {
			return h
		}
		if nums[h] > target {
			rIdx = h - 1
		} else {
			lIdx = h + 1
		}

	}
	return -1
}
