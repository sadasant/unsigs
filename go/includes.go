package unsigs

func Includes(n uint16, unsigs []uint16) bool {
	for _, u := range unsigs {
		if u == n {
			return true
		}
	}
	return false
}
