package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func readPairs(path string) [][2]uint16 {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var pairs [][2]uint16
	err = json.Unmarshal(content, &pairs)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return pairs
}

const unsigSize uint16 = 31119
const unsigSize32 uint32 = 31119

var hPairs [][2]uint16 = readPairs("../horizontal_pairs/horizontal_pairs.json")
var vPairs [][2]uint16 = readPairs("../vertical_pairs/vertical_pairs.json")
var vMap = [unsigSize][unsigSize]bool{}

func init() {
	for _, v := range vPairs {
		vMap[v[0]][v[1]] = true
	}
}

func is2x2(a, b, c, d uint16) bool {
	// Skipping horizontals since we start with them
	// return isH(Pair{a, b}) && isH(Pair{c, d}) && isV(Pair{a, c}) && isV(Pair{b, d})
	if a == c || a == d || !vMap[a][c] {
		return false
	}
	// Unique. We're assuming a,b and c,d are good horizontal pairs.
	if b == c || b == d || !vMap[b][d] {
		return false
	}
	return true
}

func main() {
	// Quick test. Both should be true.
	println(vMap[22888][28060])
	println(is2x2(10796, 10798, 10818, 10820))

	fileName := "2x2s.json"
	len_hPairs := len(hPairs)
	end := len_hPairs
	println("End:", end)

	err := ioutil.WriteFile(fileName, []byte{}, 0644)
	if err != nil {
		log.Fatal("Failed to write", fileName)
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString("["); err != nil {
		panic(err)
	}

	mar5manTop := 47379186 // 47 379 186
	startTime := time.Now()
	count := 0
	println("Running!")
L:
	for i := 0; i < end; i++ {
		matches := [][4]uint16{}
		for j := 0; j < end; j++ {
			if i == j {
				continue
			}
			h1 := hPairs[i]
			h2 := hPairs[j]
			if is2x2(h1[0], h1[1], h2[0], h2[1]) {
				matches = append(matches, [4]uint16{h1[0], h1[1], h2[0], h2[1]})
				count += 1
			}
		}
		if len(matches) == 0 {
			continue L
		}
		if count%10000 == 0 {
			seconds := int(time.Since(startTime).Seconds())
			print("\r", i, " ", seconds, " ", count, " minutes left: ", (mar5manTop*seconds)/count/60, "        ")
		}
		jsonString, _ := json.Marshal(matches)
		prefix := ","
		if i == 0 {
			prefix = ""
		}
		file.Write(append([]byte(prefix), jsonString[1:len(jsonString)-1]...))
	}

	println("Done!")
	println(count, "should be", mar5manTop)

	endString := "]"
	if _, err = file.WriteString(endString); err != nil {
		panic(err)
	}
}
