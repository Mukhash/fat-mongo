package db

import (
	"database/sql"
	"go-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTask(task *models.Task) error {
	_, err := collection.InsertOne(ctx, task)
	return err
}

func GetTasks() ([]models.Task, error) {
	filter := bson.D{{}}
	var tasks []models.Task

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var task models.Task
		err := cur.Decode(&task)
		if err != nil {
			return tasks, err
		}

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

func GetTask(id int) (*models.Task, error) {
	query := `
	SELECT id, title, body FROM tasks
	WHERE id = $1`

	row := conn.QueryRow(query, id)

	var task models.Task
	err := row.Scan(&task.ID, &task.Title, &task.Body)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *models.Task) error {
	query := `
	UPDATE tasks
	SET title = $2, body = $3
	WHERE id = $1
	`
	_, err := conn.Exec(query, task.ID, task.Title, task.Body)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(id int) error {
	query := `
	DELETE FROM tasks
	WHERE id = $1`

	_, err := conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
