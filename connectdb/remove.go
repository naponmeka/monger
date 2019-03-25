package connectdb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func removeExtractor(query string) (filter interface{}, justOne bool, err error) {
	if query == "" {
		query = "[{},{}]"
	}
	query = fmt.Sprintf("[%s]", query)
	var fields []interface{}
	err = bson.UnmarshalExtJSON([]byte(query), true, &fields)
	filter = fields[0]
	if len(fields) > 1 {
		rawOption, ok := fields[1].(map[string]interface{})
		if ok {
			if _, exists := rawOption["justOne"]; exists {
				if value, isBool := rawOption["justOne"].(bool); isBool {
					justOne = value
				}
			}
		}
	}
	return filter, justOne, err
}

func Remove(collection *mongo.Collection, filter interface{}, justOne bool) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var deleteResult *mongo.DeleteResult
	if justOne {
		deleteResult, err = collection.DeleteOne(ctx, filter, nil)
	} else {
		deleteResult, err = collection.DeleteMany(ctx, filter, nil)
	}
	var resultDoc bson.M
	if deleteResult != nil {
		resultDoc = bson.M{
			"DeletedCount": deleteResult.DeletedCount,
		}
	} else {
		resultDoc = bson.M{
			"DeletedCount": 0,
		}
	}
	results = append(results, resultDoc)
	return
}
