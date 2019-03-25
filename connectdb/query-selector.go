package connectdb

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func Query(mongoURI, db, collectionName, query string) (results []bson.M, err error) {
	collection := GetCollection(mongoURI, db, collectionName)
	if strings.HasPrefix(query, ".find(") {
		raw := GetStringInBetween(query, ".find(", ")")
		filter, option, err := findExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Find(collection, filter, option)
		if err != nil {
			return results, err
		}
	} else if strings.HasPrefix(query, ".insert(") {
		raw := GetStringInBetween(query, ".insert(", ")")
		documents, option, err := insertExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Insert(collection, documents, option)
	} else if strings.HasPrefix(query, ".update(") {
		raw := GetStringInBetween(query, ".update(", ")")
		filter, update, option, isMulti, err := updateExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Update(collection, filter, update, option, isMulti)
	} else if strings.HasPrefix(query, ".remove(") {
		raw := GetStringInBetween(query, ".remove(", ")")
		filter, justOne, err := removeExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Remove(collection, filter, justOne)
	} else {
		results = ListDocuments(mongoURI, db, collectionName)
	}
	return
}
