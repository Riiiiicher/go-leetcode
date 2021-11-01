package hash

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i = i + 1 {
		val, exist := m[target-nums[i]]
		if exist {
			return []int{val, i}
		} else {
			m[nums[i]] = i
		}
	}
	return make([]int, 0)
}
