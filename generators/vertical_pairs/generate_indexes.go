package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

func readPairs(path string) [][2]int {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var pairs [][2]int
	err = json.Unmarshal(content, &pairs)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return pairs
}

var vPairs [][2]int = readPairs("../vertical_pairs/vertical_pairs.json")
var vMap = map[int]bool{}

func toIndex(a, b int) int {
	return a*31119 + b
}

func main() {
	vMap := map[[2]int]bool{}
	for _, v := range vPairs {
		vMap[[2]int{v[0], v[1]}] = true
	}

	start := time.Now()
	println((vMap[[2]int{12142, 12397}]))
	println((vMap[[2]int{16588, 23666}]))
	println((vMap[[2]int{1507, 148}]))
	println((vMap[[2]int{144, 6}]))
	println(time.Since(start))

	booleans := [31119 * 31119]bool{}
	for _, v := range vPairs {
		booleans[toIndex(v[0], v[1])] = true
	}

	start = time.Now()
	println(booleans[toIndex(12142, 12397)])
	println(booleans[toIndex(16588, 23666)])
	println(booleans[toIndex(1507, 148)])
	println(booleans[toIndex(144, 6)])
	println(time.Since(start))

	println("Generating booleans...")

	stringBooleans := "["
	for i := 0; i < len(booleans); i++ {
		if booleans[i] {
			stringBooleans += "true,"
		} else {
			stringBooleans += "false,"
		}
	}
	result := stringBooleans[0:len(stringBooleans)-1] + "]"

	fileName := "vertical_booleans.json"
	err := ioutil.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		log.Fatal("Failed to write", fileName)
	}
}
