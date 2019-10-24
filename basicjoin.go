package piscine

func BasicJoin(strs []string) string {
	var s string
	for i := range strs {
		s = s + strs[i]
	}
	return s
}
