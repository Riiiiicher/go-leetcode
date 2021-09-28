package easy

func MaxSubArray(nums []int) int {
	rsp := nums[0]
	last := nums[0]
	for i := 1; i < len(nums); i = i + 1 {
		if last > 0 {
			last = last + nums[i]
		} else {
			last = nums[i]
		}
		if rsp < last {
			rsp = last
		}
	}
	return rsp
}
