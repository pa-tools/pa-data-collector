package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pa-tools/pa-data-collector/internal/orm"
	"github.com/pa-tools/pa-data-collector/internal/shared"
)

type postInput struct {
	shared.PaDevice
	shared.PaEntry
}

func parsePaPost(c *fiber.Ctx) error {
	var body postInput
	err := c.BodyParser(&body)
	if err != nil {
		log.Print(err.Error())
		return c.SendStatus(400)
	}
	deviceData := orm.Device{
		PaDevice: body.PaDevice,
	}
	// check if device exists, create if not
	var dbDevice orm.Device
	deviceResult := orm.Where("mac_address = ?", body.MacAddress).First(&dbDevice)
	if deviceResult.Error != nil {
		if deviceResult.Error.Error() == "record not found" {
			createResult := orm.Create(&deviceData)
			if createResult.Error != nil {
				return createResult.Error
			}
			log.Printf("Created Device record: %d\n", deviceData.ID)
			dbDevice = deviceData
		} else {
			log.Printf("Error creating device in database: %s", deviceResult.Error.Error())
			return c.SendStatus(400)
		}
	}

	entryData := orm.Entry{
		PaEntry:  body.PaEntry,
		DeviceId: dbDevice.ID,
	}
	createEntry := orm.Create(&entryData)
	if createEntry.Error != nil {
		log.Printf("Error creating device in database: %s", createEntry.Error.Error())
		return c.SendStatus(400)
	}
	log.Printf("Created Entry record: %d\n", entryData.ID)

	// write entry data and match with device
	return c.Status(200).SendString("Entry saved")
}
