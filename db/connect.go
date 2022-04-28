package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwerty"
	dbname   = "postgres"
)

var collection *mongo.Collection
var ctx = context.TODO()
var conn *mongo.Client

func CreateConnection() error {
	credentials := options.Credential{
		Username: "talgat",
		Password: "qwerty",
	}
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017/").SetAuth(credentials)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	collection = client.Database("fat-mongo").Collection("tasks")
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
