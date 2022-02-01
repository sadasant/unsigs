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
	jupiter := append(jupiter1, jupiter2[:]...)
	fmt.Printf("Jupiter's unsigs: (%v + %v) %v\n", len(jupiter1), len(jupiter2), len(jupiter))

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
	fmt.Printf("Squares made with Pawa's: %v\n", len(squares))
}
