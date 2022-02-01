// Pawa:
// > My friend Jupiter has 160 unsigs but no 2x2 blocks.
// > We want to know if by accessing my collection can he complete any 2x2 blocks.
// > Would you be able to do your calculations between 2 wallets together?
//
// Jupiter's wallets:
// - https://pool.pm/addr1qxjnqz0ay327khz9mlmdg0pgmy3zl809y0q44vlal57xqp4y480nlnjpzpwkkgfa9cxlpq8ucxt8ns6s3r7zgm06hnnsdv84r0/%400e14267a
// - https://pool.pm/addr1qyhyweywgvda8qe4mn7mwhfszux6z24usm347277vjdrxds5rgtqtftd7ztdw00w2x9806d82lqlz9jhc6k4x39sj3vqvnxapd/%400e14267a

package main

import (
	"fmt"

	unsigs "github.com/sadasant/unsigs/go"
)

func main() {
	println("Pawa's Jupiter quest")
	println("Does any unsig from Pawa's complete a 2x2 on Jupiter's?\n")

	pawa := unsigs.LoadUnsigs("pool-2022-01-31.csv")
	fmt.Printf("Pawa's unsigs: %v\n", len(pawa))

	jupiter1 := unsigs.LoadUnsigs("jupiter1-2022-01-31.csv")
	jupiter2 := unsigs.LoadUnsigs("jupiter2-2022-01-31.csv")
	jupiter3 := unsigs.LoadUnsigs("jupiter3-2022-02-01.csv")
	jupiter := append(jupiter1, jupiter2[:]...)
	jupiter = append(jupiter, jupiter3[:]...)
	jupiter = unsigs.Unique(jupiter)
	fmt.Printf("Jupiter's unsigs: (%v + %v + %v) %v\n", len(jupiter1), len(jupiter2), len(jupiter3), len(jupiter))

	jupiterHPairs := unsigs.FindHorizontalPairs(jupiter)
	fmt.Printf("Jupiter's horizontal Pairs: %v\n", len(jupiterHPairs))

	var questHPairs [][2]uint16
	for _, p := range pawa {
		for _, j := range jupiter {
			if p != j {
				if unsigs.CheckHorizontal(p, j) {
					questHPairs = append(questHPairs, [2]uint16{p, j})
					continue
				}
				if unsigs.CheckHorizontal(j, p) {
					questHPairs = append(questHPairs, [2]uint16{j, p})
					continue
				}
			}
		}
	}

	fmt.Printf("Horizontal pairs completed with 1 of Pawa's: %v\n", len(questHPairs))

	jointHPairs := append(jupiterHPairs, questHPairs[:]...)
	squares, err := unsigs.FindSquares(jointHPairs, unsigs.SquaresOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Squares made with three of Jupiter's and one of Pawa's: %v\n", len(squares))
	fmt.Printf("In details: %v\n", squares)
	pawaKeys := []uint16{}
	for _, square := range squares {
		for _, u := range square {
			for _, p := range pawa {
				if u == p {
					pawaKeys = append(pawaKeys, p)
				}
			}
		}
	}
	fmt.Printf("Made with Pawa's: %v\n", pawaKeys)

	pawaHPairs := unsigs.FindHorizontalPairs(pawa)
	var matches []struct {
		square  [4]uint16
		pawa    []uint16
		jupiter []uint16
	}
	for _, j := range jupiterHPairs {
		for _, p := range pawaHPairs {
			var matched struct {
				square  [4]uint16
				pawa    []uint16
				jupiter []uint16
			}
			matched.pawa = []uint16{p[0], p[1]}
			matched.jupiter = []uint16{j[0], j[1]}
			if unsigs.CheckSquare(p[0], p[1], j[0], j[1], unsigs.SquaresOptions{}) {
				matched.square = [4]uint16{p[0], p[1], j[0], j[1]}
				matches = append(matches, matched)
			}
			if unsigs.CheckSquare(j[0], j[1], p[0], p[1], unsigs.SquaresOptions{}) {
				matched.square = [4]uint16{j[0], j[1], p[0], p[1]}
				matches = append(matches, matched)
			}
		}
		for _, q := range questHPairs {
			var matched struct {
				square  [4]uint16
				pawa    []uint16
				jupiter []uint16
			}
			matched.jupiter = []uint16{j[0], j[1]}
			if unsigs.Includes(q[0], jupiter) {
				matched.pawa = []uint16{q[1]}
				matched.jupiter = append(matched.jupiter, q[0])
			} else {
				matched.pawa = []uint16{q[0]}
				matched.jupiter = append(matched.jupiter, q[1])
			}
			if unsigs.CheckSquare(q[0], q[1], j[0], j[1], unsigs.SquaresOptions{}) {
				matched.square = [4]uint16{q[0], q[1], j[0], j[1]}
				matches = append(matches, matched)
			}
			if unsigs.CheckSquare(j[0], j[1], q[0], q[1], unsigs.SquaresOptions{}) {
				matched.square = [4]uint16{j[0], j[1], q[0], q[1]}
				matches = append(matches, matched)
			}
		}
	}
	fmt.Printf("Squares made with two of Jupiter's and two of Pawa's: %v\n", len(matches))
	println("In detail:")
	for i, match := range matches {
		fmt.Printf("\nMatch number %v:\n", i+1)
		fmt.Printf("Square:%v\n", match.square)
		fmt.Printf("Jupiter:%v\n", match.jupiter)
		fmt.Printf("Pawa:%v\n", match.pawa)
	}
}
