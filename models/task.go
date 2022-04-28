package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Title string `bson:"title"`
	Body  string `bson:"body"`
}

type TaskId struct {
	ID    primitive.ObjectID `bson:"_id"`
	IDRaw string
	Title string `bson:"title"`
	Body  string `bson:"body"`
}
