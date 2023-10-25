package database

import (
	"sync"
)

var once sync.Once
var dbInstance *DB

type DB struct{}

func GetDB() *DB {
	once.Do(func() {
		if dbInstance == nil {
			dbInstance = &DB{}
		}
	})

	return dbInstance
}
