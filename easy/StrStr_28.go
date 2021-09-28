package easy

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i = i + 1 {
		tmp := haystack[i : i+len(needle)]
		if tmp == needle {
			return i
		}
	}
	return -1
}
