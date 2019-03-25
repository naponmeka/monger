package src

import (
	"fmt"
	"reflect"

	"github.com/globalsign/mgo/bson"
	"github.com/naponmeka/robone/connectdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createItems(mongoURI, db, collectionName, query string) (items []*TreeItem, err error) {
	documents, err := connectdb.Query(mongoURI, db, collectionName, query)
	if err != nil {
		return items, err
	}
	for _, doc := range documents {
		item := traverse(doc, doc, 0, nil)
		items = append(items, item)
	}
	return
}

func castToString(x interface{}) string {
	switch x.(type) {
	case bson.M:
		return ""
	default:
		return fmt.Sprint(x)
	}
}

func traverse(current interface{}, parent interface{}, level int, parentItem *TreeItem) *TreeItem {
	var item *TreeItem
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
	switch val := current.(type) {
	case primitive.A:
		fmt.Println(castToString(parent), fmt.Sprintf("[%d elements]", len(val)), "[array]\t\t")
		item = NewTreeItem(nil).initWith([]string{castToString(parent), fmt.Sprintf("[%d elements]", len(val)), "array"})
		level++
		for k, v := range val {
			traverse(v, fmt.Sprintf("[%d]", k), level, item)
		}
	case bson.M:
		if level == 0 {
			doc := parent.(bson.M)
			objID := doc["_id"]
			fmt.Println(castToString(objID), fmt.Sprintf("{%d fields}", len(val)), "[object]\t")
			item = NewTreeItem(nil).initWith([]string{castToString(objID), fmt.Sprintf("{%d fields}", len(val)), "object"})
		} else {
			fmt.Println(castToString(parent), fmt.Sprintf("{%d fields}", len(val)), "[object]\t")
			item = NewTreeItem(nil).initWith([]string{castToString(parent), fmt.Sprintf("{%d fields}", len(val)), "object"})
		}

		level++
		for k, v := range val {
			traverse(v, k, level, item)
		}
	default:
		fmt.Println(castToString(parent), val, fmt.Sprintf("[%s]", reflect.TypeOf(current)))
		item = NewTreeItem(nil).initWith([]string{castToString(parent), castToString(val), fmt.Sprintf("%s", reflect.TypeOf(current))})
	}
	if parentItem != nil {
		parentItem.appendChild(item)
	}
	return item
}
