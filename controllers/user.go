package controllers

import (
	"client_administration/constants"
	"client_administration/model"
	"client_administration/utils"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(c *fiber.Ctx){
	data := c.Body()
	var doc model.User

	// Unmarshal Json Data
	err := utils.Unmarshal([]byte(data) , &doc)
	if err != nil {
        fmt.Println(err)
        c.Status(400).Send("Invalid data format")
        return
    }

	// Check if required field is not empty
	validateErr := utils.CheckRequiredField(doc)
	if validateErr != nil {
		c.Status(500).Send("Required field is compulsory")
		return
	} 

	// Encrypt Password
	password := doc.Password
	hash, _ := utils.HashPassword(password) 
	doc.Password = hash
	// match := utils.CheckPasswordHash(password, "$2a$14$B1aG24XpTJbaB63MWwJuu.YV5duLIFJbw4ecPXIpDvPqIfiqSPrN.")
    // fmt.Println("Match:   ", match)

	

	usersCollection, client := model.UserModel()
	defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    }()
	insertResult, err := usersCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	userIdString, _ := insertResult.InsertedID.(primitive.ObjectID)

	registerUserResponse := constants.UserRegisterResponse{
		Email: doc.Email,
		Id:    userIdString.Hex(),
		Message: "success",
	}
	jsonResponse, err := json.Marshal(registerUserResponse)





	c.Send(jsonResponse)
}