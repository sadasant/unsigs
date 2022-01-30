package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Pairs [][2]int

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

var hPairs [][2]int = readPairs("../horizontal_pairs/horizontal_pairs.json")
var hMap = [31119 * 31119]bool{}

var vPairs [][2]int = readPairs("../vertical_pairs/vertical_pairs.json")
var vMap = [31119 * 31119]bool{}

func toIndex(a, b int) int {
	return a*31119 + b
}

func init() {
	for _, v := range vPairs {
		vMap[toIndex(v[0], v[1])] = true
	}
	for _, h := range hPairs {
		hMap[toIndex(h[0], h[1])] = true
	}
	// Garbage collecting
	hPairs = [][2]int{}
	vPairs = [][2]int{}
}

type Square [4]int

func is2x2(square Square) bool {
	a := square[0]
	b := square[1]
	c := square[2]
	d := square[3]
	// Unique. We're assuming a,b and c,d are good horizontal pairs.
	if (a == c || a == d) ||
		(b == c || b == d) {
		return false
	}
	// Skipping horizontals since we start with them
	// return isH(Pair{a, b}) && isH(Pair{c, d}) && isV(Pair{a, c}) && isV(Pair{b, d})
	return vMap[toIndex(a, c)] && vMap[toIndex(b, d)]
}

func main() {
	// Quick test. Both should be true.
	println(hMap[toIndex(10, 231)])
	println(vMap[toIndex(22888, 28060)])
	println(is2x2(Square{10796, 10798, 10818, 10820}))

	fileName := "2x2s.json"
	len_hPairs := len(hPairs)
	base := 4
	page := 1
	start := (len_hPairs / base) * (page - 1)
	end := (len_hPairs / base) * page

	if start == 0 {
		err := ioutil.WriteFile(fileName, []byte{}, 0644)
		if err != nil {
			log.Fatal("Failed to write", fileName)
		}
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString("["); err != nil {
		panic(err)
	}

	mar5manTop := 47379186
	jTemplate := "[%v,%v,%v,%v],"

	count := 0
L:
	for i := start; i < end; i++ {
		jsonString := ""
		iTemplate := ",%s"
		if i == 0 {
			iTemplate = "%s"
		}
		for j := start; j < end; j++ {
			if i == j {
				continue L
			}
			h1 := hPairs[i]
			h2 := hPairs[j]
			if is2x2(Square{h1[0], h1[1], h2[0], h2[1]}) {
				jsonString += fmt.Sprintf(jTemplate, h1[0], h1[1], h2[0], h2[1])
				if count%10000 == 0 {
					print(i, " ", j, " ", start, "-", end, " ", count, "\r")
				}
				count += 1
			}
		}
		if _, err = file.WriteString(fmt.Sprintf(iTemplate, jsonString[0:len(jsonString)-1])); err != nil {
			panic(err)
		}
	}

	println(count, "should be", mar5manTop)

	endString := "]"
	if page != base {
		endString = ","
	}
	if _, err = file.WriteString(endString); err != nil {
		panic(err)
	}
}
