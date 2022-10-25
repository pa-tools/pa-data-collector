package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	app *fiber.App
)

func initApi() {
	app = fiber.New()
	app.Use(logger.New())

	// Route Handlers
	app.Post("/post", parsePaPost)

	app.Listen(":3000")
}

type postInput struct {
	PaDevice
	PaEntry
}

func parsePaPost(c *fiber.Ctx) error {
	var body postInput
	err := c.BodyParser(&body)
	if err != nil {
		log.Print(err.Error())
		return c.SendStatus(400)
	}
	deviceData := Device{
		PaDevice: body.PaDevice,
	}
	// check if device exists, create if not
	var dbDevice Device
	deviceResult := db.Where("mac_address = ?", body.MacAddress).First(&dbDevice)
	if deviceResult.Error != nil {
		if deviceResult.Error.Error() == "record not found" {
			createResult := db.Create(&deviceData)
			if createResult.Error != nil {
				return createResult.Error
			}
			fmt.Printf("Created Device record: %d\n", deviceData.ID)
			dbDevice = deviceData
		} else {
			log.Printf("Error creating device in database: %s", deviceResult.Error.Error())
			return c.SendStatus(400)
		}
	}

	entryData := Entry{
		PaEntry:  body.PaEntry,
		DeviceId: dbDevice.ID,
	}
	createEntry := db.Create(&entryData)
	if createEntry.Error != nil {
		log.Printf("Error creating device in database: %s", createEntry.Error.Error())
		return c.SendStatus(400)
	}
	log.Printf("Created Entry record: %d\n", entryData.ID)

	// write entry data and match with device
	return c.Status(200).SendString("Entry saved")
}
