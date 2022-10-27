package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pa-tools/pa-data-collector/internal/orm"
	"github.com/samber/lo"
)

type getDevicesObj struct {
	ID           uint    `json:"id"`
	Mac          string  `json:"mac_address"`
	Name         string  `json:"name"`
	Lat          float32 `json:"latitude"`
	Lon          float32 `json:"longitude"`
	LocationType int     `json:"location_type"`
}

func getDevices(c *fiber.Ctx) error {
	devices, err := orm.GetDevices()
	if err != nil {
		return c.SendStatus(400)
	}
	deviceData := lo.Map(devices, func(item orm.Device, index int) getDevicesObj {
		locType := 1
		if strings.EqualFold(item.PaDevice.Place, "outside") {
			locType = 0
		}
		return getDevicesObj{
			ID:           item.ID,
			Mac:          item.PaDevice.MacAddress,
			Name:         item.PaDevice.Name,
			Lat:          item.PaDevice.Lat,
			Lon:          item.PaDevice.Lon,
			LocationType: locType,
		}
	})
	return c.Status(200).JSON(deviceData)
}
