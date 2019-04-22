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

func insertExtractor(query string) (documents []interface{}, option *options.InsertManyOptions, err error) {
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
	documents = fields[0].(primitive.A)
	ordered := true
	if len(fields) > 1 {
		rawOption, ok := fields[1].(primitive.D)
		if ok {
			optionMap := rawOption.Map()
			if value, isBool := optionMap["ordered"].(bool); isBool {
				ordered = value
			}
		}
	}
	option = &options.InsertManyOptions{
		Ordered: &ordered,
	}
	return documents, option, err
}

func Insert(collection *mongo.Collection, documents []interface{}, option *options.InsertManyOptions) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	rawResults, err := collection.InsertMany(ctx, documents, option)
	for _, r := range rawResults.InsertedIDs {
		bDoc := bson.M{"_id": r}
		results = append(results, bDoc)
	}
	return
}
