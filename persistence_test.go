package main

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"testing"
)

func init() {
	_ = setupDbManager()
}

func TestCursor(t *testing.T) {
	if err := dbMgr.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(radioBucket)
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var radio Radio
			_ = json.Unmarshal(v, &radio)
			//if strings.Contains(radio.Title, "搬了新家") {
			//	fmt.Printf("key=%s, value=%s\n", k, v)
			//}
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	}); nil != err {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {
	data, err := dbMgr.Find([]byte("140931"))
	if nil != err {
		t.Error(err)
		return
	} else {
		t.Log(string(data))
	}

	data, err = dbMgr.Find([]byte("1"))
	if nil != err {
		t.Error(err)
		return
	} else {
		t.Log(data)
	}
}
