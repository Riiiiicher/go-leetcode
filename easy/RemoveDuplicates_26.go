package easy

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curVal := nums[0]
	valKind := 1
	for i := 1; i < len(nums); i = i + 1 {
		if curVal != nums[i] {
			nums[valKind] = nums[i]
			valKind = valKind + 1
			curVal = nums[i]
		}
	}
	return valKind
}
