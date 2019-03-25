package connectdb

import (
	"context"
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insertExtractor(query string) (documents []interface{}, option *options.InsertManyOptions, err error) {
	if query == "" {
		query = "[{},{}]"
	}
	query = fmt.Sprintf("[%s]", query)
	var fields []interface{}
	err = bson.UnmarshalJSON([]byte(query), &fields)
	documents = fields[0].([]interface{})
	ordered := true
	if len(fields) > 1 {
		rawOption, ok := fields[1].(map[string]bool)
		if ok {
			ordered = rawOption["ordered"]
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
