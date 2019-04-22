package connectdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/naponmeka/bsonparser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func findExtractor(query string) (filter interface{}, returnOptions []*options.FindOptions, err error) {
	if query == "" {
		query = "[{}]"
	}
	query = fmt.Sprintf("[%s]", query)
	query, err = bsonparser.BsonToJson(query)
	if err != nil {
		return nil, nil, err
	}
	var fields []interface{}
	err = bson.UnmarshalExtJSON([]byte(query), true, &fields)
	if err != nil {
		return nil, nil, err
	}
	if len(fields) > 0 {
		filter = fields[0]
	}
	if len(fields) > 1 {
		option := &options.FindOptions{
			Projection: fields[1],
		}
		returnOptions = append(returnOptions, option)
	}
	return filter, returnOptions, err
}

func Find(collection *mongo.Collection, filter interface{}, option []*options.FindOptions) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := collection.Find(ctx, filter, option...)
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

func Count(collection *mongo.Collection, filter interface{}, option []*options.CountOptions) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return collection.CountDocuments(ctx, filter, option...)
}

func sortExtractor(query string) (filter interface{}, err error) {
	if query == "" {
		query = "{}"
	}
	query, err = bsonparser.BsonToJson(query)
	if err != nil {
		return nil, err
	}
	var fields interface{}
	err = bson.UnmarshalExtJSON([]byte(query), true, &fields)
	if err != nil {
		return nil, err
	}
	return fields, err
}
