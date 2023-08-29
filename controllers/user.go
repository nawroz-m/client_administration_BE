package controllers

import (
	"client_administration/constants"
	"client_administration/model"
	"client_administration/services/jwt"
	"client_administration/utils"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Register User
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


	usersCollection, client := model.UserModel()
	
	defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    }()

	// Inser User Data
	insertResult, err := usersCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	userIdString, _ := insertResult.InsertedID.(primitive.ObjectID)

	// Register User Response Data
	registerUserResponse := constants.UserRegisterResponse{
		Email: doc.Email,
		Id:    userIdString.Hex(),
		Message: "success",
	}
	jsonResponse, err := json.Marshal(registerUserResponse)
	c.Status(200).Send(jsonResponse)
}

// Login User
func LoginUser(c *fiber.Ctx){
	data := c.Body()
	var doc constants.UserLoginInfo

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

	userEmail := doc.Email
	userPassword := doc.Password
	usersCollection, client := model.UserModel()
	defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    }()

	filter := bson.D{{"email", userEmail}}


	// Find a user
	var  userInfo  model.User
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&userInfo)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	// Match the password and hashed password
	match := utils.CheckPasswordHash(userPassword, userInfo.Password)
	if match != true {
		c.Status(500).Send("Password or email address is incorrect")
		return 
	}
	userIdString:= userInfo.ID.Hex()
	
	// Generate JWT Access Token 
	token, err := jwt.GenerateJWTAccessToken(userIdString, userInfo.Email, userInfo.Password, userInfo.Role)
	if err != nil {
		c.Status(500).Send("Something went wrong")
		return
	}

	// Set a payload for user login Response
	userPayload := constants.UserLoginPayload{
		Id: userIdString,
		Email: userInfo.Email,
		Password: userInfo.Password,
		Role: userInfo.Role,
		Token: token,
	}
	jsonResponse,err := json.Marshal(userPayload)

	c.Status(200).Send(jsonResponse)
}

