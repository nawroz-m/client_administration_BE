package model

import (
	"client_administration/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Address struct {
    Street     string  `bson:"street,omitempty" validate:"required"`
    PostalCode int64  `bson:"postalcode,omitempty"`
    City       string  `bson:"city,omitempty" validate:"required"`
    Country    string  `bson:"country,omitempty" validate:"required"`
}

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	FirstName  string             `bson:"firstname,omitempty" validate:"required"`
	LastName  string             `bson:"lastname,omitempty" validate:"required"`
	Email  string             `bson:"email,omitempty" validate:"required"`
	Password  string             `bson:"password,omitempty" validate:"required"`
	telephone  int64             `bson:"telephone,omitempty" validate:"required"`
	PostalAddress  Address        `bson:"postaladdress,omitempty" validate:"required"`
	
}

func UserModel()(*mongo.Collection, *mongo.Client){
	client := config.Db();
	database := client.Database("ClientAdministration")
	usersCollection := database.Collection("users")
	return usersCollection, client
}



