package connectdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func findExtractor(query string) (filter interface{}, option *options.FindOptions, err error) {
	if query == "" {
		query = "[{}]"
	}
	query = fmt.Sprintf("[%s]", query)
	var fields []interface{}
	err = bson.UnmarshalExtJSON([]byte(query), true, &fields)
	filter = fields[0]
	if len(fields) == 2 {
		option = &options.FindOptions{
			Projection: fields[1],
		}
	}
	return filter, option, err
}

func Find(collection *mongo.Collection, filter interface{}, option *options.FindOptions) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := collection.Find(ctx, filter, option)
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
