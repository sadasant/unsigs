package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Pair [2]int
type Pairs []Pair

func readPairs(path string) Pairs {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var pairs Pairs
	err = json.Unmarshal(content, &pairs)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return pairs
}

var hPairs Pairs = readPairs("../horizontal_pairs/horizontal_pairs.json")
var hMap = map[Pair]bool{}

func isH(pair Pair) bool {
	return hMap[pair]
}

var vPairs Pairs = readPairs("../vertical_pairs/vertical_pairs.json")
var vMap = map[Pair]bool{}

func isV(pair Pair) bool {
	return vMap[pair]
}

func init() {
	for _, v := range vPairs {
		vMap[v] = true
	}
	for _, h := range hPairs {
		hMap[h] = true
	}
}

type Square [4]int

func is2x2(square Square) bool {
	unique := []int{}
L:
	for _, v := range square {
		for _, u := range unique {
			if v == u {
				continue L
			}
		}
		unique = append(unique, v)
	}
	if len(unique) != 4 {
		return false
	}
	a := square[0]
	b := square[1]
	c := square[2]
	d := square[3]
	return isH(Pair{a, b}) && isH(Pair{c, d}) && isV(Pair{a, c}) && isV(Pair{b, d})
}

func main() {
	// Quick test. Both should be true.
	println(isH(Pair{10, 231}))
	println(isV(Pair{22888, 28060}))
	println(is2x2(Square{10796, 10798, 10818, 10820}))

	// The key to make this the fastest possible is the map[Pair]bool transformation we did above.

	var squares []Square
L:
	for _, h1 := range hPairs {
		for _, h2 := range hPairs {
			if h1 == h2 {
				continue L
			}
			square := Square{h1[0], h1[1], h2[0], h2[1]}
			if is2x2(square) {
				squares = append(squares, square)
				if len(squares)%1000 == 0 {
					fmt.Printf("%v\r", len(squares))
				}
			}
		}
	}
	print(len(squares))

	file, err := json.Marshal(squares)
	if err != nil {
		log.Fatal("Failed to marshal the squares.")
	}

	fileName := "2x2s.json"
	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		log.Fatal("Failed to write", fileName)
	}
}
