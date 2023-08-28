package utils

import "encoding/json"

// Unmarshal the byte data to interface by key and value per
func Unmarshal(data []byte,  doc interface{}) error{
	err := json.Unmarshal([]byte(data), &doc)
	return err
}