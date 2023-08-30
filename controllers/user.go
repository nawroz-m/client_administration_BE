package controllers

import (
	"client_administration/aggregatepipeline"
	"client_administration/constants"
	"client_administration/model"
	"client_administration/services"
	"client_administration/services/jwt"
	"client_administration/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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
	doc.Active = true
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
	filter := bson.D{{"email", userEmail}}


	// Find a user
	var  userInfo  model.User
	responseErr := services.FindADoc(filter).Decode(&userInfo)
	if responseErr != nil {
		fmt.Print(responseErr)
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

// Update User Info
func UpdateUserInfo(c *fiber.Ctx){
	data := c.Body()
	var doc constants.UserInfoToUpdate
	// Unmarshal Json Data
	err := utils.Unmarshal([]byte(data) , &doc)
	if err != nil {
        fmt.Println(err)
        c.Status(400).Send("Invalid data format")
        return
    }

	// User information
	userData := c.Locals("user").(constants.UserLoginLocalStorage)
	objectID, err := utils.CreatObjectID(userData.Id)
    if err != nil {
        fmt.Println(err)
        c.Status(400).Send("Invalid ID format")
        return
    }
	filter := bson.D{
		{"_id",  objectID},
		// {"email",  userData.Email},
	}
    
	// Find a user
	var  userInfo  model.User
	responseErr := services.FindADoc(filter).Decode(&userInfo)
	if responseErr != nil && userInfo.Email != "" {
		fmt.Print(responseErr)
	}
	// Update for only provided fields
	update := bson.D{}
	if doc.FirstName != "" {
        update = append(update, bson.E{"firstname", doc.FirstName})
    }
    if doc.LastName != "" {
        update = append(update, bson.E{"lastname", doc.LastName})
    }
    if doc.Email != "" {
        update = append(update, bson.E{"email", doc.Email})
    }
    if doc.Telephone != 0 {
        update = append(update, bson.E{"telephone", doc.Telephone})
    }
    if doc.PostalAddress.City != "" {

        update = append(update, bson.E{"postaladdress.city", doc.PostalAddress.City})
    }
	if doc.PostalAddress.Street != "" {
        update = append(update, bson.E{"postaladdress.street", doc.PostalAddress.Street})
    }
	if doc.PostalAddress.PostalCode != 0 {
        update = append(update, bson.E{"postaladdress.postalcode", doc.PostalAddress.PostalCode})
    }
	if doc.PostalAddress.Country != "" {
        update = append(update, bson.E{"postaladdress.country", doc.PostalAddress.Country})
    }

	// Update User Info
	response := services.UpdateDocInfo(filter, update)
    c.Status(200).Send(response)
}

// Get User Info
func GetUsersInfo(c *fiber.Ctx){
	id := c.Query("id")
	search := c.Query("search")
	fmt.Println(id, search)
	
	// User information
	userData := c.Locals("user").(constants.UserLoginLocalStorage)
	objectID, err := utils.CreatObjectID(userData.Id)
    if err != nil {
        c.Status(400).Send("Invalid ID")
        return
    }
	filter := bson.D{
		{"_id",  objectID},
		// {"email",  userData.Email},
	}
    
	// Find a user
	var  userInfo  model.User
	responseErr := services.FindADoc(filter).Decode(&userInfo)
	if responseErr != nil && userInfo.Email != "" {
		fmt.Print(responseErr)
	}
	if userInfo.Role != "addmin" {
		 c.Status(400).Send("Not accessible, pleas check your credentioal")
        return
	}
	searchFilter := constants.SearchUserData{
		Id: id,
		Search: search,
		Active: true,
	}

	// Get user 
	pipeline := aggregatepipeline.GetUserDataForAdminPipeline(searchFilter)
	usersCollection, client := model.UserModel()
	defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            return
        }
    }()

	cursor, err := usersCollection.Aggregate(context.TODO(), pipeline)

	if err != nil {
    // handle error
		fmt.Println("Something wrong while aggregating")
    	return
	}

	var users []map[string]interface{}
	if err := cursor.All(context.TODO(), &users); err != nil {
		// handle error
		fmt.Println("Something wrong while cursing errr")
		return
	}
	jsonData, err := json.Marshal(users)

    c.Status(http.StatusOK).Send(jsonData)
}