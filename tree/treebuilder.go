package tree

import (
	"fmt"
	"reflect"

	"github.com/naponmeka/robone/connectdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateItems(
	mongoURI string,
	db string,
	collectionName string,
	query string,
	skip int,
	limit int,
	timeout int,
) (items []*TreeItem, documents []bson.M, err error) {
	documents, err = connectdb.Query(mongoURI, db, collectionName, query, skip, limit, timeout)
	if err != nil {
		return nil, nil, err
	}
	for _, doc := range documents {
		item := Traverse(doc, doc, 0, nil)
		items = append(items, item)
	}
	return
}

func CastToString(x interface{}) string {
	switch x.(type) {
	case bson.M:
		return ""
	default:
		return fmt.Sprint(x)
	}
}

func Traverse(current interface{}, parent interface{}, level int, parentItem *TreeItem) *TreeItem {
	var item *TreeItem
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
	switch val := current.(type) {
	case primitive.A:
		fmt.Println(CastToString(parent), fmt.Sprintf("[%d elements]", len(val)), "[array]\t\t")
		item = NewTreeItem(nil).initWith([]string{CastToString(parent), fmt.Sprintf("[%d elements]", len(val)), "array"})
		level++
		for k, v := range val {
			Traverse(v, fmt.Sprintf("[%d]", k), level, item)
		}
	case bson.M:
		if level == 0 {
			doc := parent.(bson.M)
			objID := doc["_id"]
			fmt.Println(CastToString(objID), fmt.Sprintf("{%d fields}", len(val)), "[object]\t")
			item = NewTreeItem(nil).initWith([]string{CastToString(objID), fmt.Sprintf("{%d fields}", len(val)), "object"})
		} else {
			fmt.Println(CastToString(parent), fmt.Sprintf("{%d fields}", len(val)), "[object]\t")
			item = NewTreeItem(nil).initWith([]string{CastToString(parent), fmt.Sprintf("{%d fields}", len(val)), "object"})
		}

		level++
		for k, v := range val {
			Traverse(v, k, level, item)
		}
	default:
		fmt.Println(CastToString(parent), val, fmt.Sprintf("[%s]", reflect.TypeOf(current)))
		item = NewTreeItem(nil).initWith([]string{CastToString(parent), CastToString(val), fmt.Sprintf("%s", reflect.TypeOf(current))})
	}
	if parentItem != nil {
		parentItem.appendChild(item)
	}
	return item
}
