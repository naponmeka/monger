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

func ListDB(mongoURI string) []string {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	results, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func ListCollection(mongoURI string, db string) (results []string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017/admin"))
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

// func ListDocuments(mongoURI, db, collection string) {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017/admin"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	cur, err := client.Database(db).Collection(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cur.Close(ctx)
// 	for cur.Next(ctx) {
// 		elem := CollectionName{}
// 		if err := cur.Decode(&elem); err != nil {
// 			log.Fatal(err)
// 		}
// 		if elem.Name != "system.js" {
// 			results = append(results, elem.Name)
// 		}
// 	}
// 	return
// }
