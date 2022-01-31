package unsigs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func LoadUnsigs(path string) []uint16 {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	r := regexp.MustCompile(`unsig\d+`)
	matches := r.FindAllString(string(content), -1)

	appended := map[uint64]bool{}

	var unsigs []uint16
	for _, str := range matches {
		n, _ := strconv.ParseUint(str[len("unsig"):], 10, 16)
		if appended[n] {
			continue
		}
		appended[n] = true
		unsigs = append(unsigs, uint16(n))
	}
	return unsigs
}

func LoadPairs(path string) [][2]uint16 {
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
