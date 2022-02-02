package main

import (
	"fmt"

	unsigs "github.com/sadasant/unsigs/go"
)

func main() {
	println("Results for Pawa (before Jan 31)")

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

	found2x3s, err := unsigs.Find2x3s(vPairs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("2x3s: %v\n", len(found2x3s))

	found3x2s, err := unsigs.Find3x2s(hPairs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("3x2s: %v\n", len(found3x2s))

	println("\nResults for Pawa (after Jan 31)")

	pool = unsigs.LoadUnsigs("pool-2022-01-31.csv")
	fmt.Printf("Unsigs: %v\n", len(pool))
	unsigs.WriteJson("pawa.json", pool)

	vPairs = unsigs.FindVerticalPairs(pool)
	fmt.Printf("Vertical Pairs: %v\n", len(vPairs))
	unsigs.WriteJson("pawa_vpairs.json", vPairs)

	hPairs = unsigs.FindHorizontalPairs(pool)
	fmt.Printf("Horizontal Pairs: %v\n", len(hPairs))
	unsigs.WriteJson("pawa_hpairs.json", hPairs)

	squares, err = unsigs.FindSquares(hPairs, unsigs.SquaresOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Squares: %v\n", len(squares))
	unsigs.WriteJson("pawa_squares.json", squares)
}
