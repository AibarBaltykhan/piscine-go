package piscine

func Join(strs []string, sep string) string {
	var s string
	k := 0
	for i := range strs {
		i = i
		k++
	}
	for j := 0; j < k-1; j++ {
		s = s + strs[j] + sep
	}
	s = s + strs[k-1]
	return s
}
