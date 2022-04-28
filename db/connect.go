package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host     = "localhost"
	port     = 27017
	user     = "talgat"
	password = "qwerty"
	dbname   = "fat-mongo"
)

var collection *mongo.Collection
var ctx = context.Background()
var conn *mongo.Client

func CreateConnection() error {
	clientOptions := options.Client()

	clientOptions.ApplyURI("mongodb://localhost:27017/")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	collection = client.Database("fat-mongo").Collection("tasks")
	mod := mongo.IndexModel{
		Keys: bson.M{
			"title": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	collection.Indexes().CreateOne(ctx, mod)

	// testTask := models.Task{Title: "Home", Body: "Wash dishes"}
	// _, err = collection.InsertOne(ctx, testTask)
	// if err != nil {
	// 	return err
	// }

	conn = client
	return nil
}

func GetDB() *mongo.Client {
	if conn != nil {
		return conn
	}
	return nil
}

func GetCollection() *mongo.Collection {
	if collection != nil {
		return collection
	}
	return nil
}
