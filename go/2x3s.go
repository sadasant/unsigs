package unsigs

import "errors"

// Shape:
// 0 1 2
// 3 4 5
func Check2x3(x [6]uint16) bool {
	if len(Unique(x[:])) != 6 {
		return false
	}
	vertical := CheckVertical(x[0], x[3]) && CheckVertical(x[1], x[4]) && CheckVertical(x[2], x[5])
	horizontal := CheckHorizontal(x[0], x[1]) && CheckHorizontal(x[1], x[2]) && CheckHorizontal(x[3], x[4]) && CheckHorizontal(x[4], x[5])
	return vertical && horizontal
}

func Find2x3s(verticalPairs [][2]uint16) ([][6]uint16, error) {
	matches := [][6]uint16{}
	found := map[[6]uint16]bool{}
	for _, pair1 := range verticalPairs {
		if !CheckVertical(pair1[0], pair1[1]) {
			return nil, errors.New("Find2x3s only supports passing vertical pairs")
		}
		for _, pair2 := range verticalPairs {
			if pair1 == pair2 {
				continue
			}
			for _, pair3 := range verticalPairs {
				if pair1 == pair3 || pair2 == pair3 {
					continue
				}
				shape2x3 := [6]uint16{pair1[0], pair2[0], pair3[0], pair1[1], pair2[1], pair3[1]}
				if Check2x3(shape2x3) && !found[shape2x3] {
					matches = append(matches, shape2x3)
					found[shape2x3] = true
				}
			}
		}
	}
	return matches, nil
}
