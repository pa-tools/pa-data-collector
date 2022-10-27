package orm

import "gorm.io/gorm"

func Create(value interface{}) (tx *gorm.DB) {
	if db == nil {
		InitDb()
	}
	return db.Create(value)
}

func Where(query interface{}, args ...interface{}) (tx *gorm.DB) {
	if db == nil {
		InitDb()
	}
	return db.Where(query, args)
}
