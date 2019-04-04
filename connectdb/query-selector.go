package connectdb

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Query(
	mongoURI string,
	db string,
	collectionName string,
	query string,
	skip int,
	limit int,
) (results []bson.M, err error) {
	query = strings.TrimSpace(query)
	collection := GetCollection(mongoURI, db, collectionName)
	if strings.HasPrefix(query, ".find(") {
		if strings.HasSuffix(query, ".count()") {
			raw := GetStringInBetween(strings.TrimSuffix(query, ".count()"), ".find(", ")")
			filter, _, _ := findExtractor(raw)
			c, err := Count(collection, filter, nil)
			results = append(results, bson.M{"count": c})
			return results, err
		} else {
			raw := GetStringInBetween(query, ".find(", ")")
			filter, fOptions, err := findExtractor(raw)
			if skip >= 0 {
				fOptions = append(fOptions, &options.FindOptions{
					Skip: &[]int64{int64(skip)}[0],
				})
			}
			if limit >= 0 {
				fOptions = append(fOptions, &options.FindOptions{
					Limit: &[]int64{int64(limit)}[0],
				})

			}
			if err != nil {
				return results, err
			}
			results, err = Find(collection, filter, fOptions)
			if err != nil {
				return results, err
			}
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
	} else if strings.HasPrefix(query, ".aggregate(") {
		raw := GetStringInBetween(query, ".aggregate(", ")")
		steps, err := aggregateExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Aggregate(collection, steps)
	} else {
		results = ListDocuments(mongoURI, db, collectionName)
	}
	return
}
