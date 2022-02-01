package unsigs

import (
	"encoding/json"
	"io/ioutil"
)

func WriteJson(fileName string, contents interface{}) error {
	result, err := json.Marshal(contents)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, result, 0644)
}
