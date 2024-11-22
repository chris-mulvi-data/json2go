package internal

import (
	"encoding/json"
	"errors"
	"os"
)

// ReadFile reads a JSON file and will unmarshal it into a map[string]interface{}
func ReadFile(path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		// return nil, errors.New("failed to read the json file")
		return nil, err
	}

	isValidJson := json.Valid(content)
	if !isValidJson {
		return nil, errors.New("invalid json")
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
