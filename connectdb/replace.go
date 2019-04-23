package connectdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/naponmeka/bsonparser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func replaceExtractor(query string) (filter interface{}, update interface{}, option *options.ReplaceOptions, err error) {
	if query == "" {
		query = "[{},{},{}]"
	}
	query = fmt.Sprintf("[%s]", query)
	query, err = bsonparser.BsonToJson(query)
	if err != nil {
		return nil, nil, nil, err
	}
	var fields []interface{}
	err = bson.UnmarshalExtJSON([]byte(query), true, &fields)
	if len(fields) < 2 {
		return nil, nil, nil, errors.New("filter & replacement required")
	}
	filter = fields[0]
	update = fields[1]
	option = &options.ReplaceOptions{}
	if len(fields) > 2 {
		rawOption, ok := fields[2].(map[string]interface{})
		if ok {
			if _, exists := rawOption["upsert"]; exists {
				if value, isBool := rawOption["upsert"].(bool); isBool {
					option.Upsert = &value
				}
			}
		}
	}
	return filter, update, option, err
}

func Replace(timeout int, collection *mongo.Collection, filter interface{}, replacement interface{}, option *options.ReplaceOptions) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	var updateResult *mongo.UpdateResult
	updateResult, _ = collection.ReplaceOne(ctx, filter, replacement)
	var resultDoc bson.M
	if updateResult != nil {
		resultDoc = bson.M{
			"MatchedCount":  updateResult.MatchedCount,
			"ModifiedCount": updateResult.ModifiedCount,
			"UpsertedCount": updateResult.UpsertedCount,
			"UpsertedID":    updateResult.UpsertedID,
		}
	} else {
		resultDoc = bson.M{
			"MatchedCount":  0,
			"ModifiedCount": 0,
			"UpsertedCount": 0,
			"UpsertedID":    "-",
		}
	}
	results = append(results, resultDoc)
	return
}
