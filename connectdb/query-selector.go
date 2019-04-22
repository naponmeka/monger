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
			query = strings.TrimSuffix(query, ".count()")
			if i := strings.Index(query, ".sort("); i > 0 {
				query = query[:i]
			}
			raw := GetStringInBetween(query, ".find(", ")")
			filter, _, _ := findExtractor(raw)
			c, err := Count(collection, filter, nil)
			results = append(results, bson.M{"count": c})
			return results, err
		} else {
			sortRaw := ""
			if i := strings.Index(query, ".sort("); i > 0 {
				sortRaw = GetStringInBetween(query, ".sort(", ")")
				query = query[:i]
			}
			raw := GetStringInBetween(query, ".find(", ")")
			filter, fOptions, err := findExtractor(raw)
			if err != nil {
				return results, err
			}
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
			sort, err := sortExtractor(sortRaw)
			if err != nil {
				return results, err
			}
			fOptions = append(fOptions, &options.FindOptions{
				Sort: sort,
			})
			results, err = Find(collection, filter, fOptions)
		}

	} else if strings.HasPrefix(query, ".insert(") {
		raw := GetStringInBetween(query, ".insert(", ")")
		documents, option, err := insertExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Insert(collection, documents, option)
	} else if strings.HasPrefix(query, ".insertMany(") {
		raw := GetStringInBetween(query, ".insertMany(", ")")
		documents, option, err := insertExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Insert(collection, documents, option)
	} else if strings.HasPrefix(query, ".insertOne(") {
		raw := GetStringInBetween(query, ".insertOne(", ")")
		document, option, err := insertOneExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = InsertOne(collection, document, option)
	} else if strings.HasPrefix(query, ".update(") {
		raw := GetStringInBetween(query, ".update(", ")")
		filter, update, option, isMulti, err := updateExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Update(collection, filter, update, option, isMulti)
	} else if strings.HasPrefix(query, ".updateMany(") {
		raw := GetStringInBetween(query, ".updateMany(", ")")
		filter, update, option, isMulti, err := updateExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Update(collection, filter, update, option, isMulti)
	} else if strings.HasPrefix(query, ".updateOne(") {
		raw := GetStringInBetween(query, ".updateOne(", ")")
		filter, update, option, _, err := updateExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Update(collection, filter, update, option, false)
	} else if strings.HasPrefix(query, ".remove(") {
		raw := GetStringInBetween(query, ".remove(", ")")
		filter, justOne, err := removeExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Remove(collection, filter, justOne)
	} else if strings.HasPrefix(query, ".deleteOne(") {
		raw := GetStringInBetween(query, ".deleteOne(", ")")
		filter, _, err := removeExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Remove(collection, filter, true)
	} else if strings.HasPrefix(query, ".deleteMany(") {
		raw := GetStringInBetween(query, ".deleteMany(", ")")
		filter, _, err := removeExtractor(raw)
		if err != nil {
			return results, err
		}
		results, err = Remove(collection, filter, false)
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
