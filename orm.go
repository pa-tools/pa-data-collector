package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Device struct {
	gorm.Model
	PaDevice PaDevice `gorm:"embedded"`
	Entries  []Entry
}

type Entry struct {
	gorm.Model
	PaEntry  PaEntry `gorm:"embedded"`
	DeviceId uint
}

func initDb() {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Entry{})
	db.AutoMigrate(&Device{})
}
