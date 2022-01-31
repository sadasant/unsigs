package main

import (
	"fmt"

	unsigs "github.com/sadasant/unsigs/go"
)

func main() {
	println("Results for $sadasant")
	pool := unsigs.LoadUnsigs("pool-2022-01-31.csv")
	fmt.Printf("Unsigs: %v\n", pool)
	vPairs := unsigs.FindVerticalPairs(pool)
	fmt.Printf("Vertical Pairs: %v\n", vPairs)
	hPairs := unsigs.FindHorizontalPairs(pool)
	fmt.Printf("Horizontal Pairs: %v\n", hPairs)
}
