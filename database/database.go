package database

import "sync"

var once sync.Once
var dbInstance *database

type database struct{}

func GetInstance() *database {
	once.Do(func() {
		if dbInstance == nil {
			dbInstance = &database{}
		}
	})

	return dbInstance
}
