package connectdb

import (
	"context"
	"log"
	"time"

	"github.com/naponmeka/bsonparser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func aggregateExtractor(query string) (steps interface{}, err error) {
	if query == "" {
		query = "[]"
	}
	query, err = bsonparser.BsonToJson(query)
	err = bson.UnmarshalExtJSON([]byte(query), true, &steps)
	return steps, nil
}

func Aggregate(collection *mongo.Collection, steps interface{}) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := collection.Aggregate(ctx, steps)
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
