package connectdb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ListIndex(timeout int, collection *mongo.Collection) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	// defer cancel()
	cur, err := collection.Indexes().List(ctx)
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
