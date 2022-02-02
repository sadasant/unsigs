package unsigs

// Shape:
// 0 1 2
// 3 4 5
// 6 7 8
func Check3x3(x [9]uint16) bool {
	if len(Unique(x[:])) != 9 {
		return false
	}
	parameter2x3 := [6]uint16{}
	copy(parameter2x3[:], x[0:6])
	a2x3 := Check2x3(parameter2x3)

	copy(parameter2x3[:], x[3:9])
	b2x3 := Check2x3(parameter2x3)
	return a2x3 && b2x3
}

func Find3x3s(pairs [][2]uint16) ([][9]uint16, error) {
	matches := [][9]uint16{}
	found := map[[9]uint16]bool{}
	for _, pair1 := range pairs {
		for _, pair2 := range pairs {
			if pair1 == pair2 {
				continue
			}
			for _, pair3 := range pairs {
				if pair1 == pair3 || pair2 == pair3 {
					continue
				}
				for _, pair4 := range pairs {
					if pair1 == pair4 || pair2 == pair4 || pair3 == pair4 {
						continue
					}
					for _, pair5 := range pairs {
						if pair1 == pair5 || pair2 == pair5 || pair3 == pair5 || pair4 == pair5 {
							continue
						}
						for _, pair6 := range pairs {
							if pair1 == pair6 || pair2 == pair6 || pair3 == pair6 || pair4 == pair6 || pair5 == pair6 {
								continue
							}
							if pair1[1] != pair2[0] || pair3[1] != pair4[0] || pair5[1] != pair6[0] {
								continue
							}
							shape3x3 := [9]uint16{pair1[0], pair1[1], pair2[1], pair3[0], pair3[1], pair4[1], pair5[0], pair5[1], pair6[1]}
							if Check3x3(shape3x3) && !found[shape3x3] {
								matches = append(matches, shape3x3)
								found[shape3x3] = true
							}
						}
					}
				}
			}
		}
	}
	return matches, nil
}
