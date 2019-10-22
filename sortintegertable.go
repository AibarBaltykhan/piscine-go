
package piscine

func SortIntegerTable(table []int) {
	length := 0
	c := 0
	for _, v := range table {
		v = v
		length++
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if table[i] > table[j] {
				c = table[i]
				table[i] = table[j]
				table[j] = c
			}
		}
	}
}