package easy

func RemoveElement(nums []int, val int) int {
	lenth := len(nums)
	for i := 0; i < len(nums); {
		if nums[i] == val {
			lenth = lenth - 1
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i = i + 1
		}
	}
	return lenth
}
