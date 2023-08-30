package utils

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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



func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func CreatObjectID(Id string)(primitive.ObjectID, error){
	objectID, err := primitive.ObjectIDFromHex(Id)
	return objectID, err
}
