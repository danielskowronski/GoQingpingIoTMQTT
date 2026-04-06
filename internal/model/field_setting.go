// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

// SettingField... contain read/write fields, used to set parameters and to report settings back

const (
	// HEX - those are used to set parameters and to report settings back
	SettingFieldIntervalUploadingHEX        int16 = 0x04 // uint16, minutes
	SettingFieldIntervalRecordingHEX        int16 = 0x05 // uint16, seconds
	SettingFieldReportTemperatureHigher     int16 = EvtReportTemperatureHigher
	SettingFieldReportTemperatureLower      int16 = EvtReportTemperatureLower
	SettingFieldReportHumidityHigher        int16 = EvtReportHumidityHigher
	SettingFieldReportHumidityLower         int16 = EvtReportHumidityLower
	SettingFieldReportPressureHigher        int16 = EvtReportPressureHigher
	SettingFieldReportPressureLower         int16 = EvtReportPressureLower
	SettingFieldReportTemperatureAuxHigher  int16 = EvtReportTemperatureAuxHigher
	SettingFieldReportTemperatureAuxLower   int16 = EvtReportTemperatureAuxLower
	SettingFieldUnitOfTemperatureHEX        int16 = 0x19 // uint16, 0 for Celsius, 1 for Fahrenheit
	SettingFieldIntervalEventRepeat         int16 = 0x1B // uint16, seconds
	SettingFieldOffsetTempHumiMain          int16 = 0x2F // 2x uint16, step is 0.1, first 2 bytes are temperature, last 2 bytes are humidity
	SettingFieldOffsetTempHumiExt           int16 = 0x30 // 2x uint16, step is 0.1, first 2 bytes are temperature, last 2 bytes are humidity
	SettingFieldOffsetPressure              int16 = 0x31 // uint16, step is 0.1
	SettingFieldReportingDelay              int16 = 0x37 // uint16, seconds
	SettingFieldReportCO2Higher             int16 = EvtReportCO2Higher
	SettingFieldReportCO2Lower              int16 = EvtReportCO2Lower
	SettingFieldIntervalRecordingCO2HEX     int16 = 0x3B // uint16, minutes
	SettingFieldCO2Ranges                   int16 = 0x3C // 2x uint16, 2 bytes for low range, 2 bytes for high range
	SettingFieldAutoShutdownOnBattery       int16 = 0x3D // uint16, minutes
	SettingFieldCO2AutoCalibration          int16 = 0x40 // uint16, 0 for disabled, 1 for enabled
	SettingFieldOffsetCO2Percentage         int16 = 0x3F // uint16, step is 0.1%
	SettingFieldOffsetCO2Value              int16 = 0x45 // uint16, step is 1 ppm
	SettingFieldOffsetTemperatureValue      int16 = 0x46 // uint16, step is 0.1 degrees
	SettingFieldOffsetTemperaturePercentage int16 = 0x47 // uint16, step is 0.1%
	SettingFieldOffsetHumidityValue         int16 = 0x48 // uint16, step is 0.1 %RH
	SettingFieldOffsetHumidityPercentage    int16 = 0x49 // uint16, step is 0.1%
	SettingFieldOffsetPM25Value             int16 = 0x4B // uint16, step is 1 ppm
	SettingFieldOffsetPM25Percentage        int16 = 0x4C // uint16, step is 0.1%
	SettingFieldOffsetPM10Value             int16 = 0x4D // uint16, step is 1 ppm
	SettingFieldOffsetPM10Percentage        int16 = 0x4E // uint16, step is 0.1%
	SettingFieldTemperatureRanges           int16 = 0x4F // 2x uint16, 2 bytes for low range, 2 bytes for high range
	SettingFieldHumidityRanges              int16 = 0x50 // 2x uint16, 2 bytes for low range, 2 bytes for high range
	SettingFieldPM25Ranges                  int16 = 0x51 // 2x uint16, 2 bytes for low range, 2 bytes for high range
	SettingFieldPM10Ranges                  int16 = 0x52 // 2x uint16, 2 bytes for low range, 2 bytes for high range
	SettingFieldTVOCRanges                  int16 = 0x53 // 2x uint16, 2 bytes for low range, 2 bytes for high range
	SettingFieldNoiseRanges                 int16 = 0x54 // 2x uint16, 2 bytes for low range, 2 bytes for high range
	SettingFieldLightRanges                 int16 = 0x55 // 2x uint16, 2 bytes for low range, 2 bytes for high range
	SettingFieldReportTVOCHigher            int16 = EvtReportTVOCHigher
	SettingFieldReportTVOCLower             int16 = EvtReportTVOCLower
	SettingFieldReportPM25Higher            int16 = EvtReportPM25Higher
	SettingFieldReportPM25Lower             int16 = EvtReportPM25Lower
	SettingFieldReportPM10Higher            int16 = EvtReportPM10Higher
	SettingFieldReportPM10Lower             int16 = EvtReportPM10Lower
	SettingFieldReportNoiseHigher           int16 = EvtReportNoiseHigher
	SettingFieldReportNoiseLower            int16 = EvtReportNoiseLower
	SettingFieldReportLightHigher           int16 = EvtReportLightHigher
	SettingFieldReportLightLower            int16 = EvtReportLightLower

	// JSON
	SettingFieldSettingsJSON             string = "settings"              // structure
	SettingFieldIntervalUploadingJSON    string = "report_interval"       // (under SettingFieldSettingsJSON) int
	SettingFieldIntervalRecordingJSON    string = "collect_interval"      // (under SettingFieldSettingsJSON) int
	SettingFieldIntervalRecordingCO2JSON string = "co2_sampling_interval" // (under SettingFieldSettingsJSON) int
	SettingFieldIntervalRecordingPM      string = "pm_sampling_interval"  // (under SettingFieldSettingsJSON) int
	SettingFieldUnitOfTemperatureJSON    string = "temperature_unit"      // (under SettingFieldSettingsJSON) char, "C" or "F"
	SettingFieldNightModeStart           string = "night_mode_start_time" // (under SettingFieldSettingsJSON) int, minutes since midnight
	SettingFieldNightModeEnd             string = "night_mode_end_time"   // (under SettingFieldSettingsJSON) int, minutes since midnight
)
