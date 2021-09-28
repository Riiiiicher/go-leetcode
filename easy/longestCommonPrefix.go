package easy

func LongestCommonPrefix(strs []string) string {
	rsp := ""
	for pos := 0; ; pos = pos + 1 {
		tmp := ""
		for i := 0; i < len(strs); i = i + 1 {
			if len(strs[i]) <= pos {
				return rsp
			}
			if tmp == "" {
				tmp = strs[i][pos : pos+1]
			} else {
				if tmp != strs[i][pos:pos+1] {
					return rsp
				}
			}
		}
		rsp = rsp + tmp
	}
}
