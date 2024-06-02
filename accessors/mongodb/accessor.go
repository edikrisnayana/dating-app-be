package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Accessor struct {
	Collection *mongo.Collection
	Context    context.Context
}

func CreateAccessor(hostname string, dbName string, collectionName string) *Accessor {
	mongoUri := fmt.Sprintf("mongodb://%s:27017", hostname)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri).SetReadPreference(readpref.Secondary()))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return &Accessor{Collection: client.Database(dbName).Collection(collectionName), Context: context.Background()}
}
