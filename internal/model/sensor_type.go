package model

type SensorTypeInfo struct {
	Unit          string
	HADeviceClass string // see https://www.home-assistant.io/integrations/sensor
	Description   string
}

type SensorType string

const (
	SensorTypeUnknown      SensorType = "unknown"
	SensorBattery          SensorType = "battery"
	SensorTemperature      SensorType = "temperature"
	SensorTemperatureAux   SensorType = "temperature_aux"
	SensorPressure         SensorType = "pressure"
	SensorHumidity         SensorType = "humidity"
	SensorHumidityAux      SensorType = "humidity_aux"
	SensorPM1              SensorType = "pm1"
	SensorPM25             SensorType = "pm25"
	SensorPM10             SensorType = "pm10"
	SensorCO2              SensorType = "co2"
	SensorCO2Concentration SensorType = "co2_percentage"
	SensorVOC              SensorType = "voc"
	SensorNoise            SensorType = "noise"
	SensorLight            SensorType = "light"
	SensorTVOC             SensorType = "tvoc"
	SensorRadon            SensorType = "radon"
	SensorSignalStrength   SensorType = "signal_strength"
)

var SensorTypeInfoMap = map[SensorType]SensorTypeInfo{
	// default units from spec, to be checked if settings override them
	SensorTypeUnknown:      {Unit: "", HADeviceClass: "None", Description: "Unknown sensor type"},
	SensorBattery:          {Unit: "%", HADeviceClass: "battery", Description: "Battery level"},
	SensorTemperature:      {Unit: "°C", HADeviceClass: "temperature", Description: "Temperature"},
	SensorTemperatureAux:   {Unit: "°C", HADeviceClass: "temperature", Description: "AUX temperature"},
	SensorPressure:         {Unit: "hPa", HADeviceClass: "pressure", Description: "Atmospheric pressure"}, // this is not native unit, required conversion from raw
	SensorHumidity:         {Unit: "%", HADeviceClass: "humidity", Description: "Humidity"},
	SensorHumidityAux:      {Unit: "%", HADeviceClass: "humidity", Description: "AUX humidity"},
	SensorPM1:              {Unit: "µg/m³", HADeviceClass: "pm1", Description: "Particulate matter 1.0"},
	SensorPM25:             {Unit: "µg/m³", HADeviceClass: "pm25", Description: "Particulate matter 2.5"},
	SensorPM10:             {Unit: "µg/m³", HADeviceClass: "pm10", Description: "Particulate matter 10"},
	SensorCO2:              {Unit: "ppm", HADeviceClass: "carbon_dioxide", Description: "CO2"},
	SensorCO2Concentration: {Unit: "%", HADeviceClass: "None", Description: "CO2 concentration"},
	SensorVOC:              {Unit: "index", HADeviceClass: "volatile_organic_compounds_parts", Description: "Volatile organic compounds"},
	SensorNoise:            {Unit: "dB", HADeviceClass: "sound_pressure", Description: "Noise level"},
	SensorLight:            {Unit: "lux", HADeviceClass: "illuminance", Description: "Light level"},
	SensorTVOC:             {Unit: "ppb", HADeviceClass: "volatile_organic_compounds", Description: "Total volatile organic compounds"},
	SensorRadon:            {Unit: "index", HADeviceClass: "None", Description: "Radon level"},
	SensorSignalStrength:   {Unit: "dBm", HADeviceClass: "signal_strength", Description: "Wi-Fi signal strength"},
}
