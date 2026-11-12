package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const DBNAME = "infus"

func ConnectDatabase(configs Configuration) *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(configs.MongoUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	return client
}

func FetchData[T any](client *mongo.Client, collection string, filter bson.D, opts options.Lister[options.FindOptions]) ([]T, error) {
	col := client.Database(DBNAME).Collection(collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := col.Find(ctx, filter, opts)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	var result []T
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func InsertOne[T any](client *mongo.Client, collection string, insertData T) (*mongo.InsertOneResult, error) {
	col := client.Database(DBNAME).Collection(collection)

	result, err := col.InsertOne(context.TODO(), insertData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
