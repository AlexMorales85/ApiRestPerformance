package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var test int

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}

func GetTasksEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")

	var tasks []Task

	collection := client.Database("tasks").Collection("tasks")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message : "` + err.Error() + `" }`))
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message : "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(tasks)
}

func main() {

	fmt.Println("Starting application...")

	test = 10

	client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetTasksEndpoint).Methods("GET")

	http.ListenAndServe("0.0.0.0:4000", router)
}
