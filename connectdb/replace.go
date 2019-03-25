package connectdb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Replace(collection *mongo.Collection, filter interface{}, replacement interface{}, option *options.ReplaceOptions) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
