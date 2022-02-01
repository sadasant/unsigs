package unsigs

import (
	"errors"
)

type SquaresOptions struct {
	VerticalOrientation bool
}

// Shape:
// 1 2
// 3 4
// Shape if abVertical:
// 1 3
// 2 4
func CheckSquare(a uint16, b uint16, c uint16, d uint16, options SquaresOptions) bool {
	check := CheckVertical
	if options.VerticalOrientation {
		check = CheckHorizontal
	}
	// We're assuming a,b and c,d are good horizontal pairs.
	if a == c || a == d || !check(a, c) {
		return false
	}
	if b == c || b == d || !check(b, d) {
		return false
	}
	return true
}

func FindSquares(pairs [][2]uint16, options SquaresOptions) ([][4]uint16, error) {
	if options.VerticalOrientation {
		if !CheckVertical(pairs[0][0], pairs[0][1]) {
			return nil, errors.New("if `vertical` is set, the `pairs` array should be composed of vertical pairs")
		}
	} else {
		if !CheckHorizontal(pairs[0][0], pairs[0][1]) {
			return nil, errors.New("if `vertical` is not set, the `pairs` array should be composed of horizontal pairs")
		}
	}
	matches := [][4]uint16{}
	for _, pair1 := range pairs {
		for _, pair2 := range pairs {
			if pair1 == pair2 {
				continue
			}
			if CheckSquare(pair1[0], pair1[1], pair2[0], pair2[1], options) {
				matches = append(matches, [4]uint16{pair1[0], pair1[1], pair2[0], pair2[1]})
			}
		}
	}
	return matches, nil
}
