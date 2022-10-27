package orm

import (
	"github.com/pa-tools/pa-data-collector/internal/shared"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Device struct {
	gorm.Model
	PaDevice shared.PaDevice `gorm:"embedded"`
	Entries  []Entry
}

type Entry struct {
	gorm.Model
	PaEntry  shared.PaEntry `gorm:"embedded"`
	DeviceId uint
}

func InitDb() {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Entry{})
	db.AutoMigrate(&Device{})
}
