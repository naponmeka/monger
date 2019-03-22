package connectdb

import (
	"context"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type CollectionName struct {
	Name string `json:"name"`
}

func ListDB(mongoURI string) (results []string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	rawResults, err := client.ListDatabaseNames(ctx, bson.M{})
	for _, r := range rawResults {
		if r != "admin" && r != "config" && r != "local" {
			results = append(results, r)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ListCollection(mongoURI string, db string) (results []string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	cur, err := client.Database(db).ListCollections(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		elem := CollectionName{}
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		if elem.Name != "system.js" {
			results = append(results, elem.Name)
		}
	}
	return
}

func ListDocuments(mongoURI, db, collectionName string) (results []bson.M) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(db).Collection(collectionName)
	if err != nil {
		log.Fatal(err)
	}
	cur, err := collection.Find(ctx, bson.M{})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		elem := &bson.M{}
		if err := cur.Decode(elem); err != nil {
			log.Fatal(err)
		}
		results = append(results, *elem)
	}
	return
}
