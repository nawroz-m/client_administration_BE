package services

import (
	"client_administration/model"
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Find A Single Document
func FindADoc(filter interface{})(*mongo.SingleResult){
	usersCollection, client := model.UserModel()
	defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            return
        }
    }()

	// Find a user
	Response := usersCollection.FindOne(context.TODO(), filter)
	fmt.Println("typ of response: ", reflect.TypeOf(Response))
	return Response
}

// Update A Document Field
func UpdateDocInfo( filter primitive.D, update primitive.D)string{
	usersCollection, client := model.UserModel()
	defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            return
        }
    }()	

	if len(update) > 0 {
        updateDoc := bson.D{bson.E{"$set", update}}
		// Update document
        _, err := usersCollection.UpdateOne(context.TODO(), filter, updateDoc)
        if err != nil {
            return "Update failed"
        }
        return "Data Updated "
    } else {
		return "No fields to update"
    }
	
}


