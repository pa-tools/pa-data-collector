package orm

import (
	"github.com/pa-tools/pa-data-collector/internal/shared"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	PaEntry  shared.PaEntry `gorm:"embedded"`
	DeviceId uint
}

func CreateEntry(entry Entry) (Entry, error) {
	transaction := db.Create(&entry)
	return entry, transaction.Error
}
