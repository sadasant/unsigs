package unsigs

import "errors"

// Shape:
// 0 1
// 2 3
// 4 5
func Check3x2(x [6]uint16) bool {
	if len(Unique(x[:])) != 6 {
		return false
	}
	horizontal := CheckHorizontal(x[0], x[1]) && CheckHorizontal(x[2], x[3]) && CheckHorizontal(x[4], x[5])
	vertical := CheckVertical(x[0], x[2]) && CheckVertical(x[2], x[4]) && CheckVertical(x[1], x[3]) && CheckVertical(x[3], x[5])
	return vertical && horizontal
}

func Find3x2s(horizontalPairs [][2]uint16) ([][6]uint16, error) {
	matches := [][6]uint16{}
	found := map[[6]uint16]bool{}
	for _, pair1 := range horizontalPairs {
		if !CheckHorizontal(pair1[0], pair1[1]) {
			return nil, errors.New("Find3x2s only supports passing horizontal pairs")
		}
		for _, pair2 := range horizontalPairs {
			if pair1 == pair2 {
				continue
			}
			for _, pair3 := range horizontalPairs {
				if pair1 == pair3 || pair2 == pair3 {
					continue
				}
				shape3x2 := [6]uint16{pair1[0], pair1[1], pair2[0], pair2[1], pair3[0], pair3[1]}
				if Check3x2(shape3x2) && !found[shape3x2] {
					matches = append(matches, shape3x2)
					found[shape3x2] = true
				}
			}
		}
	}
	return matches, nil
}
