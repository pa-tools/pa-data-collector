package main

type PaDevice struct {
	MacAddress string  `gorm:"index:idx_mac,unique" json:"SensorId"`
	Name       string  `json:"Geo"`
	Lat        float32 `json:"lat"`
	Lon        float32 `json:"lon"`
	Place      string  `json:"place"`
}

type PaEntry struct {
	EntryTimestamp string  `json:"DateTime"`
	Temperature    int     `json:"current_temp_f"`
	Humidity       int     `json:"current_humidity"`
	Dewpoint       int     `json:"current_dewpoint_f"`
	Pressure       float32 `json:"pressure"`
}
