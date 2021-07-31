package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var mongo_err error

func ConnectToDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB

	Client, mongo_err = mongo.Connect(context.TODO(), clientOptions)

	if mongo_err != nil {
		log.Fatal(mongo_err)
	}

	fmt.Println("Connected to MongoDB!")
}

func ServerStarted() {
	fmt.Println("servers started")
	data, err := Client.Database("grave_yard").Collection("Jonathan").Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(data.Decode(bson.D{}))

}
