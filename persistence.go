package main

import (
	bolt "go.etcd.io/bbolt"
	"sync"
)

var (
	radioBucket = []byte("radio")
	dbMgr       *dbManager
)

type dbManager struct {
	sync.RWMutex
	db *bolt.DB
}

func setupDbManager(refresh bool) error {
	db, err := bolt.Open("record.db", 0600, nil)
	if nil != err {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		if refresh {
			if err = tx.DeleteBucket(radioBucket); nil != err {
				return err
			}
		}
		_, err := tx.CreateBucketIfNotExists(radioBucket)
		return err
	})
	if nil != err {
		return err
	}
	dbMgr = &dbManager{
		db: db,
	}
	return nil
}

func (m *dbManager) Find(key []byte) ([]byte, error) {
	m.RLock()
	defer m.RUnlock()
	var result []byte
	if err := m.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(radioBucket)
		bytes := bucket.Get(key)
		result = bytes
		return nil
	}); nil != err {
		return nil, err
	}

	return result, nil
}

func (m *dbManager) All() ([][]byte, error) {
	m.RLock()
	defer m.RUnlock()
	var result [][]byte
	err := m.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(radioBucket)
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			result = append(result, v)
		}
		return nil
	})
	return result, err
}

func (m *dbManager) Insert(key, value []byte) error {
	m.Lock()
	defer m.Unlock()
	err := m.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(radioBucket)
		if data := bucket.Get(key); len(data) <= 0 {
			return bucket.Put(key, value)
		}
		return nil
	})
	return err
}
