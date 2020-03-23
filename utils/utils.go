package utils

func SplitMultiBlank(s string) []string {
	var res []string
	i, j := 0, 0
	for i < len(s) && j < len(s) {
		if s[i] == ' ' {
			i++
			j++
		}
		if j < len(s) && s[j] != ' ' {
			j++
			if j >= len(s) {
				res = append(res, s[i:j])
				break
			}
		} else if j > i {
			res = append(res, s[i:j])
			j++
			i = j
		}
	}
	return res
}
