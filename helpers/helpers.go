package helpers

import (
	"context"
	"log"

	"github.com/manuelrobert/test_1_server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneStudent(student models.Student, collection *mongo.Collection) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func UpdateOneStudent(studentId string, student models.Student, collection *mongo.Collection) *mongo.UpdateResult {
	result, err := collection.UpdateOne(context.Background(), bson.M{"student_id": studentId}, bson.M{"$set": student})
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetOneStudent(studentId string, collection *mongo.Collection) models.Student {
	var student models.Student

	err := collection.FindOne(context.Background(), bson.M{"student_id": studentId}).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}
	return student
}

func DeleteOneStudent(studentId string, collection *mongo.Collection) *mongo.DeleteResult {
	result, err := collection.DeleteOne(context.Background(), bson.M{"student_id": studentId})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func GetManyStudents(collection *mongo.Collection) []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var students []primitive.M
	for cursor.Next(context.Background()) {
		var student bson.M
		err := cursor.Decode(&student)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	defer cursor.Close(context.Background())
	return students
}
