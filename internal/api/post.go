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

	dbDevice, err := orm.GetOrCreateDevice(orm.Device{
		PaDevice: body.PaDevice,
	}, "mac_address = ?", body.MacAddress)
	if err != nil {
		return c.SendStatus(400)
	}

	createEntry, err := orm.CreateEntry(orm.Entry{
		PaEntry:  body.PaEntry,
		DeviceId: dbDevice.ID,
	})
	if err != nil {
		log.Printf("Error creating device in database: %s", err.Error())
		return c.SendStatus(400)
	}
	log.Printf("Created Entry record: %d\n", createEntry.ID)

	return c.Status(200).SendString("Entry saved")
}
