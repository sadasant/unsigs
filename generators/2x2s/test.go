package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func readJson(path string) [][4]uint16 {
	println(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var squares [][4]uint16
	err = json.Unmarshal(content, &squares)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return squares
}

func main() {
	var squares0 [][4]uint16 = readJson("./2x2s_0.json")
	var squares1 [][4]uint16 = readJson("./2x2s_1.json")
	var squares2 [][4]uint16 = readJson("./2x2s_2.json")
	var squares3 [][4]uint16 = readJson("./2x2s_3.json")
	var squares4 [][4]uint16 = readJson("./2x2s_4.json")

	println("Lengths...")
	println("0:", len(squares0))
	println("1:", len(squares1))
	println("2:", len(squares2))
	println("3:", len(squares3))
	println("4:", len(squares4))

	squares := [][4]uint16{}
	squares = append(squares, squares0[:]...)
	squares = append(squares, squares1[:]...)
	squares = append(squares, squares2[:]...)
	squares = append(squares, squares3[:]...)
	squares = append(squares, squares4[:]...)

	mar5manTop := 47379186 // 47 379 186

	// Note: I'm seeing 38 673 215, not 47 379 186. Why?
	println(len(squares), "should be", mar5manTop)
}
