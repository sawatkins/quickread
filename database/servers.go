package database

import (
	"github.com/sawatkins/quickread/models"

	"go.mongodb.org/mongo-driver/bson"
)

// add a server to the database
func AddServer(server models.Server) (interface{}, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Convert the Server struct to a BSON document
	data, err := bson.Marshal(server)
	if err != nil {
		return nil, err
	}

	// Get the collection from the database
	collection := client.Database(dbName).Collection("servers")

	// Insert the data into the collection
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

// get a server from the database
func GetServer(id string) (models.Server, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Get the collection from the database
	collection := client.Database(dbName).Collection("servers")

	// Find the document
	filter := bson.M{"id": id}
	var result models.Server
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.Server{}, err
	}

	return result, nil
}

// get all servers from the database
func GetAllServers() ([]models.Server, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Get the collection from the database
	collection := client.Database(dbName).Collection("servers")

	// Find the document
	filter := bson.M{}
	var result []models.Server
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// update a server in the database
func UpdateServer(id string, server models.Server) (interface{}, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Convert the Server struct to a BSON document
	data, err := bson.Marshal(server)
	if err != nil {
		return nil, err
	}

	// Get the collection from the database
	collection := client.Database(dbName).Collection("servers")

	// Update the document
	filter := bson.M{"id": id}
	result, err := collection.ReplaceOne(ctx, filter, data)
	// result, err := collection.UpdateOne(ctx, filter, data)
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}

// delete a server from the database
func DeleteServer(id string) (interface{}, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Get the collection from the database
	collection := client.Database(dbName).Collection("servers")

	// Delete the document
	filter := bson.M{"id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

// get all active public servers from the database
func GetActivePublicServers() ([]models.Server, error) {
	client, ctx, dbName, cancel := Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	// Get the collection from the database
	collection := client.Database(dbName).Collection("servers")

	// Find the document
	filter := bson.M{"public": true, "status": "active"}
	var result []models.Server
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
