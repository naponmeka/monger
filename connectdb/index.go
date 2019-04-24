package connectdb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ListIndex(timeout int, collection *mongo.Collection) (results []bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	// defer cancel()
	cur, err := collection.Indexes().List(ctx)
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

func CreateIndex(mongoURI string,
	db string,
	collectionName string,
	name string,
	keys bson.M,
	sparse bool,
	unique bool,
	inBackground bool,
	expireAfterSeconds int,
) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	indexOptions := options.IndexOptions{}
	indexOptions.SetName(name)
	indexOptions.SetSparse(sparse)
	indexOptions.SetUnique(unique)
	indexOptions.SetBackground(inBackground)
	if expireAfterSeconds > 0 {
		indexOptions.SetExpireAfterSeconds(int32(expireAfterSeconds))
	}

	index := mongo.IndexModel{
		Keys:    keys,
		Options: &indexOptions,
	}
	_, err = client.Database(db).Collection(collectionName).Indexes().CreateOne(ctx, index)
	if err != nil {
		return
	}
	return nil
}

func DropIndex(mongoURI, db, collectionName, name string) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	_, err = client.Database(db).Collection(collectionName).Indexes().DropOne(ctx, name)
	return
}
