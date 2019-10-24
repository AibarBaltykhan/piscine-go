package piscine

func BasicJoin(strs []string) string {
	var s string
	for i := range strs {
		s = append(s, strs[i])
	}
	return s
}
