package main

import (
	"fmt"

	unsigs "github.com/sadasant/unsigs/go"
)

func main() {
	println("Results for ArtAddict69 (before June 2)")

	pool := unsigs.LoadUnsigs("pool.csv")
	fmt.Printf("Unsigs: %v\n", len(pool))

	vPairs := unsigs.FindVerticalPairs(pool)
	fmt.Printf("Vertical Pairs: %v\n", len(vPairs))

	hPairs := unsigs.FindHorizontalPairs(pool)
	fmt.Printf("Horizontal Pairs: %v\n", len(hPairs))

	squares, err := unsigs.FindSquares(hPairs, unsigs.SquaresOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Squares: %v\n", len(squares))
	unsigs.WriteJson("squares.json", squares)

	// No 2x3s
	// found2x3s, err := unsigs.Find2x3s(vPairs)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("2x3s: %v\n", len(found2x3s))

	// No 3x2s
	// found3x2s, err := unsigs.Find3x2s(hPairs)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("3x2s: %v\n", len(found3x2s))

	println("\nProposal")

	newUnsigs := []uint16{12397, 15490, 26642, 30223, 31022, 12142, 12683, 15032}
	proposal := unsigs.Concat(pool, newUnsigs)
	proposal = unsigs.Exclude(proposal, []uint16{1376})
	fmt.Printf("Proposal total: %v\n", len(proposal))

	vPairs = unsigs.FindVerticalPairs(proposal)
	fmt.Printf("Vertical Pairs: %v\n", len(vPairs))
	unsigs.WriteJson("proposal_vpairs.json", vPairs)

	hPairs = unsigs.FindHorizontalPairs(proposal)
	fmt.Printf("Horizontal Pairs: %v\n", len(hPairs))
	unsigs.WriteJson("proposal_hpairs.json", hPairs)

	squares, err = unsigs.FindSquares(hPairs, unsigs.SquaresOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Squares: %v\n", len(squares))
	unsigs.WriteJson("proposal_squares.json", squares)
}
