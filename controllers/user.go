package controllers

import (
	"client_administration/model"
	"client_administration/utils"
	"context"
	"fmt"

	"github.com/gofiber/fiber"
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
	// password := doc.Password



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

	c.Send(insertResult.InsertedID)
}