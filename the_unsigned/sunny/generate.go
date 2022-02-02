package main

import (
	"fmt"

	unsigs "github.com/sadasant/unsigs/go"
)

func main() {
	println("\nResults for Sunny (2022-02-01)")

	pool := unsigs.LoadUnsigs("pool-2022-02-01.csv")
	fmt.Printf("Unsigs: %v\n", len(pool))
	unsigs.WriteJson("sunny.json", pool)

	vPairs := unsigs.FindVerticalPairs(pool)
	fmt.Printf("Vertical Pairs: %v\n", len(vPairs))
	unsigs.WriteJson("vpairs.json", vPairs)

	hPairs := unsigs.FindHorizontalPairs(pool)
	fmt.Printf("Horizontal Pairs: %v\n", len(hPairs))
	unsigs.WriteJson("hpairs.json", hPairs)

	squares, err := unsigs.FindSquares(hPairs, unsigs.SquaresOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Squares: %v\n", len(squares))
	unsigs.WriteJson("squares.json", squares)
}
