package unsigs

func Unique(unsigs []uint16) []uint16 {
	uniqueMap := map[uint16]bool{}
	result := []uint16{}
	for _, u := range unsigs {
		if !uniqueMap[u] {
			result = append(result, u)
			uniqueMap[u] = true
		}
	}
	return result
}
