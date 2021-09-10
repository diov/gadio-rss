package main

import (
	bolt "go.etcd.io/bbolt"
	"sync"
)

var (
	radioBucket = []byte("radio")
	mgr         *dbManager
)

type dbManager struct {
	sync.RWMutex
	db *bolt.DB
}

func setupManager() error {
	db, err := bolt.Open("record.db", 0600, nil)
	if nil != err {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(radioBucket)
		return err
	})
	if nil != err {
		return err
	}
	mgr = &dbManager{
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

func (m *dbManager) Insert(key, value []byte) error {
	m.Lock()
	defer m.Unlock()
	err := m.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(radioBucket)
		return bucket.Put(key, value)
	})
	return err
}
