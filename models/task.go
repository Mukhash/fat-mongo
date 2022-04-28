package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title"`
	Body  string             `bson:"body"`
}
