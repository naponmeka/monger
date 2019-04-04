package connectdb

import (
	"context"
	"fmt"
	"time"

	"github.com/naponmeka/bsonparser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func updateExtractor(query string) (filter interface{}, update interface{}, option *options.UpdateOptions, isMulti bool, err error) {
	if query == "" {
		query = "[{},{},{}]"
	}
	query = fmt.Sprintf("[%s]", query)
	query, err = bsonparser.BsonToJson(query)
	var fields []interface{}
	err = bson.UnmarshalExtJSON([]byte(query), true, &fields)
	filter = fields[0]
	update = fields[1]
	option = &options.UpdateOptions{}
	if len(fields) > 2 {
		rawOption, ok := fields[2].(map[string]interface{})
		if ok {
			if _, exists := rawOption["upsert"]; exists {
				if value, isBool := rawOption["upsert"].(bool); isBool {
					option.Upsert = &value
				}
			}
			if _, exists := rawOption["multi"]; exists {
				if value, isBool := rawOption["multi"].(bool); isBool {
					isMulti = value
				}
			}
		}
	}
	return filter, update, option, isMulti, err
}

func Update(collection *mongo.Collection, filter interface{}, update interface{}, option *options.UpdateOptions, isMulti bool) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var updateResult *mongo.UpdateResult
	if isMulti {
		updateResult, err = collection.UpdateMany(ctx, filter, update, option)
	} else {
		updateResult, err = collection.UpdateOne(ctx, filter, update, option)
	}
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
