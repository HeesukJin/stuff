package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoDBConnect() *mongo.Client {
	MongoDBClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017/"))
    if err != nil {
        panic(err)
    }

	return MongoDBClient
}