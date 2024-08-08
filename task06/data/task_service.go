package data

import (
	"context"
	"os"
	"errors"
	"time"
	"log"
	models "task06/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"fmt"
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

var taskCollection = GetClient().Collection("tasks")
var userCollection = GetClient().Collection("users")


func RegisterUser(newUser models.User) error {

	unhashedPwd := newUser.Password 
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(unhashedPwd), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error while hashing the password")
	}

	newUser.Password = string(hashedPwd)

	_, err = userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return errors.New("error while registering user")
	}
	return nil
}

func AuthenticateUser(userInfo models.User) (string, error) {

	var user models.User
	userEmail := userInfo.Email
	filter := bson.D{{Key : "email", Value : userEmail}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return "",errors.New("invalid credentials")
	}

	hashedPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userInfo.Password))
	
	if err != nil {
		return "", errors.New("invalid credentials")
	}


	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id" : user.ID,
		"user_email": user.Email,
		"exp" : time.Now().Add(time.Hour*24),
	})

	jwtToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", errors.New("error while generating token")
	}

	return jwtToken, nil

}

func FindAllTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	cur, err := taskCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return tasks, errors.New("error while fetching data")
	}

	for cur.Next(context.TODO()) {
		var elem models.Task

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

func AddTask(newTask models.Task) error {
	_, err := taskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return errors.New("there was an error while inserting task to the DB")
	}
	return nil
}

func SearchByID(id string) (models.Task, error) {
	filter := bson.D{{Key : "id", Value: id}}
	var foundTask models.Task

	err := taskCollection.FindOne(context.TODO(), filter).Decode(&foundTask)
	fmt.Println(err)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return foundTask, nil
}

func DeleteByID(id string) (error) {
	filter := bson.D{{Key : "id", Value: id}}
	_, err := taskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("document Not Found")
	}
	return nil
}
func ModifyTask(modified models.Task, id string) error {
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

	_, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("document not found")
	}
	return nil
}

