package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Replace the uri string with your MongoDB deployment's connection string.
const uri = "mongodb+srv://<user>:<password>@<cluster-url>?retryWrites=true&w=majority"

func main() {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// begin findOneUsageExample
	coll := client.Database("sample_mflix").Collection("movies")
	var result bson.D
	err = coll.FindOne(ctx, bson.D{{"title", "The Room"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
	// end findOneUsageExample

	relaxedJSONBytes, err := bson.MarshalExtJSON(result, false, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", relaxedJSONBytes)
}
