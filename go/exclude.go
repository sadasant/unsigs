package unsigs

func Exclude(a []uint16, b []uint16) []uint16 {
	result := []uint16{}
Loop:
	for _, u := range a {
		for _, u2 := range b {
			if u == u2 {
				continue Loop
			}
		}
		result = append(result, u)
	}
	return result
}
