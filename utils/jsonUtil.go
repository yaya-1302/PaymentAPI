package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFromFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v", filePath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(v)
	if err != nil {
		return fmt.Errorf("could not decode JSON from file %s: %v", filePath, err)
	}

	return nil
}

func WriteToFile(filePath string, v interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file %s: %v", filePath, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(v)
	if err != nil {
		return fmt.Errorf("could not encode JSON to file %s: %v", filePath, err)
	}

	return nil
}
