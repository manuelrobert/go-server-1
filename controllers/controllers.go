package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manuelrobert/test_1_server/helpers"
	"github.com/manuelrobert/test_1_server/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "<mongodb-connection-string>"
const dbName = "testdb"
const colName = "student"

var collection *mongo.Collection

func init() {
	fmt.Println("Database connection initiating...")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection successfull")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance ready")
}

func CreateOneStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}

	result := helpers.InsertOneStudent(student, collection)
	json.NewEncoder(w).Encode(result)
}

func UpdateOneStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r)
	result := helpers.UpdateOneStudent(params["id"], student, collection)
	json.NewEncoder(w).Encode(result)
}

func GetOneStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	result := helpers.GetOneStudent(params["id"], collection)
	json.NewEncoder(w).Encode(result)
}

func DeleteOneStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result := helpers.DeleteOneStudent(params["id"], collection)
	json.NewEncoder(w).Encode(result)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	result := helpers.GetManyStudents(collection)
	json.NewEncoder(w).Encode(result)
}
