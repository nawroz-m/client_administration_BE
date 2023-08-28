package utils

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// Unmarshal the byte data to interface by key and value per
func Unmarshal(data []byte,  doc interface{}) error{
	err := json.Unmarshal([]byte(data), &doc)
	return err
}

// Check the required field
func CheckRequiredField(doc interface {}) error{

	validate := validator.New()
	validateErr := validate.Struct(doc)
	return validateErr
}