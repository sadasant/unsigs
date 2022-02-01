package unsigs

var HPairs [][2]uint16 = LoadPairs("../../generators/horizontal_pairs/horizontal_pairs.json")
var VPairs [][2]uint16 = LoadPairs("../../generators/vertical_pairs/vertical_pairs.json")

var hMap = [UnsigSize][UnsigSize]bool{}
var vMap = [UnsigSize][UnsigSize]bool{}

func init() {
	println("Preparing pairs...")
	for _, v := range VPairs {
		vMap[v[0]][v[1]] = true
	}
	for _, h := range HPairs {
		hMap[h[0]][h[1]] = true
	}
}

func Combinations(unsigs []uint16) [][2]uint16 {
	var pairs [][2]uint16
	for i, u1 := range unsigs {
		for j, u2 := range unsigs {
			if i != j {
				pairs = append(pairs, [2]uint16{u1, u2})
			}
		}
	}
	return pairs
}

func CheckVertical(a, b uint16) bool {
	return vMap[a][b]
}

func FindVerticalPairs(unsigs []uint16) [][2]uint16 {
	pairs := Combinations(unsigs)
	var result [][2]uint16
	for _, pair := range pairs {
		if CheckVertical(pair[0], pair[1]) {
			result = append(result, pair)
		}
	}
	return result
}

func CheckHorizontal(a, b uint16) bool {
	return hMap[a][b]
}

func FindHorizontalPairs(unsigs []uint16) [][2]uint16 {
	pairs := Combinations(unsigs)
	var result [][2]uint16
	for _, pair := range pairs {
		if CheckHorizontal(pair[0], pair[1]) {
			result = append(result, pair)
		}
	}
	return result
}
