package shared

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
	Pm2_5Aqi       int     `json:"pm2.5_aqi"`
	Pm2_5AqiA      int     `json:"pm2.5_aqi_a"`
	Pm2_5AqiB      int     `json:"pm2.5_aqi_b"`
	Pm2_5Cf1       float32 `json:"pm2_5_cf_1"`
	Pm2_5Cf1A      float32 `json:"pm2_5_cf_1_a"`
	Pm2_5Cf1B      float32 `json:"pm2_5_cf_1_b"`
	Pm2_5Atm       float32 `json:"pm2_5_atm"`
	Pm2_5AtmA      float32 `json:"pm2_5_atm_a"`
	Pm2_5AtmB      float32 `json:"pm2_5_atm_b"`
	Pm10Cf1        float32 `json:"pm1_0_cf_1"`
	Pm10Cf1A       float32 `json:"pm1_0_cf_1_a"`
	Pm10Cf1B       float32 `json:"pm1_0_cf_1_b"`
	Pm10Atm        float32 `json:"pm1_0_atm"`
	Pm10AtmA       float32 `json:"pm1_0_atm_a"`
	Pm10AtmB       float32 `json:"pm1_0_atm_b"`
}
