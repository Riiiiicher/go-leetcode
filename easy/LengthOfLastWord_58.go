package easy

func lengthOfLastWord(s string) int {
	rsp := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			rsp++
		} else if rsp > 0 {
			return rsp
		}
	}
	return rsp
}
