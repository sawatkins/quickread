package database

import (
	"github.com/sawatkins/upfast.tf-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

// add a user to the database
func AddUser(user models.User) (interface{}, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Convert the User struct to a BSON document
	data, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}

	// Get the collection from the database
	collection := client.Database(dbName).Collection("users")

	// Insert the data into the collection
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

// get a user from the database
func GetUser(id string) (models.User, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Get the collection from the database
	collection := client.Database(dbName).Collection("users")

	// Find the document
	filter := bson.M{"id": id}
	var result models.User
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.User{}, err
	}

	return result, nil
}

// get all users from the database
// generated with copilot, not sure if it works
func GetAllUsers() ([]models.User, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Get the collection from the database
	collection := client.Database(dbName).Collection("users")

	// Find the document
	filter := bson.M{}
	var result []models.User
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return []models.User{}, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return []models.User{}, err
		}
		result = append(result, user)
	}
	if err := cur.Err(); err != nil {
		return []models.User{}, err
	}

	return result, nil
}

// delete a user from the database
func DeleteUser(id string) (int64, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Get the collection from the database
	collection := client.Database(dbName).Collection("users")

	// Delete the document
	filter := bson.M{"id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}