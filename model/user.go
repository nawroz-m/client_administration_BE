package model

import (
	"client_administration/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type Contact struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Email string             `bson:"email,omitempty"`
	Tags  []string           `bson:"tags,omitempty"`
}

func UserModel()(*mongo.Collection, *mongo.Client){
	client := config.Db();
	database := client.Database("ClientAdministration")
	usersCollection := database.Collection("users")
	return usersCollection, client
}



