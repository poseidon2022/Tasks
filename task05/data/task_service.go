package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	task "task05/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error while connecting to the Database")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Database connecction fatally failed")
	}

	db := client.Database("task_management")
	fmt.Println("Database connection setup")
	return db
}

var collection = GetClient().Collection("tasks")

func FindAllTasks() ([]*task.Task, error) {
	var tasks []*task.Task
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return tasks, errors.New("error while fetching data")
	}

	for cur.Next(context.TODO()) {
		var elem task.Task

		err := cur.Decode(&elem)
		if err != nil {
			return tasks, errors.New("error while fetching data")
		}
		tasks = append(tasks, &elem)
	}

	err = cur.Err()
	if err != nil {
		return tasks, errors.New("error while fetching data")
	}

	cur.Close(context.TODO())

	return tasks, nil
}

func AddTask(newTask task.Task) error {
	_, err := collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return errors.New("there was an error while inserting task to the DB")
	}
	return nil
}

func SearchByID(id string) (task.Task, error) {
	filter := bson.D{{Key : "id", Value: id}}
	var foundTask task.Task

	err := collection.FindOne(context.TODO(), filter).Decode(&foundTask)
	fmt.Println(err)
	if err != nil {
		return task.Task{}, errors.New("task not found")
	}
	return foundTask, nil
}

func DeleteByID(id string) (error) {
	filter := bson.D{{Key : "id", Value: id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("document Not Found")
	}
	return nil
}
func ModifyTask(modified task.Task, id string) error {
	filter := bson.D{{Key : "id", Value : id}}

	update := bson.D{}

	if modified.Title != "" {
		update = append(update, bson.E{
			Key : "$set", Value : bson.D{{
				Key : "title", Value : modified.Title,
			}}})
	}

	if modified.Status != "" {
		update = append(update, bson.E{
			Key : "$set", Value : bson.D{{
				Key : "status", Value : modified.Status,
			}}})
	}

	if modified.Description != "" {
		update = append(update, bson.E{
			Key : "$set", Value : bson.D{{
				Key : "description", Value : modified.Description,
			}}})
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("document not found")
	}
	return nil
}

