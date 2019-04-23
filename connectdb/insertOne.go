package connectdb

import (
	"context"
	"fmt"
	"time"

	"github.com/naponmeka/bsonparser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insertOneExtractor(query string) (document interface{}, option *options.InsertOneOptions, err error) {
	if query == "" {
		query = "[[],{}]"
	}
	query = fmt.Sprintf("[%s]", query)
	query, err = bsonparser.BsonToJson(query)
	if err != nil {
		return nil, nil, err
	}
	var fields []interface{}
	err = bson.UnmarshalExtJSON([]byte(query), true, &fields)
	document = fields[0].(primitive.D)
	option = &options.InsertOneOptions{}
	return document, option, err
}

func InsertOne(timeout int, collection *mongo.Collection, document interface{}, option *options.InsertOneOptions) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	insertOneResult, err := collection.InsertOne(ctx, document, option)
	if err != nil {
		return results, err
	}
	bDoc := bson.M{"_id": insertOneResult.InsertedID}
	results = append(results, bDoc)
	return results, nil
}
