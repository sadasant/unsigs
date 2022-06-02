package unsigs

func Concat(a []uint16, b []uint16) []uint16 {
	uniqueMap := map[uint16]bool{}
	result := []uint16{}
	for _, u := range a {
		if !uniqueMap[u] {
			result = append(result, u)
			uniqueMap[u] = true
		}
	}
	for _, u := range b {
		if !uniqueMap[u] {
			result = append(result, u)
			uniqueMap[u] = true
		}
	}
	return result
}
