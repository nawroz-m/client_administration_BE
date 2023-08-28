package controllers

import (
	"client_administration/model"
	"client_administration/utils"
	"context"
	"fmt"

	"github.com/gofiber/fiber"
)

func InserUserInfo(c *fiber.Ctx){
	data := c.Body()
	var doc model.Contact
	err := utils.Unmarshal([]byte(data) , &doc)
	if err != nil {
        fmt.Println(err)
        c.Status(400).Send("Invalid data format")
        return
    }
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