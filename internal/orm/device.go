package orm

import (
	"log"

	"github.com/pa-tools/pa-data-collector/internal/shared"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	PaDevice shared.PaDevice `gorm:"embedded"`
	Entries  []Entry
}

func GetOrCreateDevice(deviceData Device, query interface{}, args ...interface{}) (Device, error) {
	deviceResult := db.Where(query, args).First(&deviceData)
	if deviceResult.Error != nil {
		if deviceResult.Error.Error() == "record not found" {
			createResult := db.Create(&deviceData)
			if createResult.Error == nil {
				log.Printf("Created Device record: %d\n", deviceData.ID)
			}
			return deviceData, createResult.Error
		} else {
			log.Printf("Error creating device in database: %s\n", deviceResult.Error.Error())
		}
	} else {
		log.Printf("Found existing device: %d\n", deviceData.ID)
	}
	return deviceData, deviceResult.Error
}

func GetDevices() ([]Device, error) {
	var devices []Device
	transaction := db.Find(&devices)
	if transaction.Error != nil {
		log.Printf("Error getting all devices: %s\n", transaction.Error.Error())
	}
	return devices, transaction.Error
}
