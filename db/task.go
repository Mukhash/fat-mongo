package db

import (
	"errors"
	"fmt"
	"go-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTask(task *models.Task) error {
	_, err := collection.InsertOne(ctx, task)
	return err
}

func GetTasks() ([]models.TaskId, error) {
	filter := bson.D{{}}
	var tasks []models.TaskId

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var task models.TaskId
		err := cur.Decode(&task)
		if err != nil {
			return tasks, err
		}

		task.IDRaw = task.ID.Hex()
		tasks = append(tasks, task)
	}

	if err = cur.Err(); err != nil {
		return tasks, nil
	}

	cur.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}
	return tasks, nil
}

func GetTask(id string) (*models.TaskId, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s: %s\n", "IdFromHex", err.Error()))
	}

	filter := bson.M{"_id": objID}
	res := collection.FindOne(ctx, filter)

	var task models.TaskId

	err = res.Decode(&task)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s: %s\n", "Decode", err.Error()))
	}

	task.IDRaw = id
	return &task, nil
}

func UpdateTask(task *models.TaskId) error {
	objID, err := primitive.ObjectIDFromHex(task.IDRaw)
	if err != nil {
		return errors.New(fmt.Sprintf("%s: %s\n", "IdFromHex", err.Error()))
	}

	_, err = collection.UpdateByID(ctx, objID, bson.D{
		{"$set", bson.D{
			{"title", task.Title},
			{"body", task.Body},
		}},
	})

	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
