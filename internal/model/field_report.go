// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

// ReportField... contain read-only fields

const (
	// HEX - those are used by device in reports
	ReportFieldDeviceID             int16 = 0x01 // string
	ReportFieldDeviceSN             int16 = 0x02 // string
	ReportFieldHistoricalDataSix    int16 = 0x03 // length, then series of 6-byte records
	ReportFieldHistoricalDataEight  int16 = 0x33 // length, then series of 8-byte records (incl. AUX sensor)
	ReportFieldSIMCardNumber        int16 = 0x16 // string
	ReportFieldAllDataSent          int16 = 0x17 // uint8, when going to standby, 0 indicates more data will come, 1 means all data is sent
	ReportFieldHardwareVersion      int16 = 0x22 // string
	ReportFieldNextConnection       int16 = 0x24 // timestamp of next connection
	ReportFieldUSBPluggedIn         int16 = 0x2C // uint8, 0 for no, 1 for yes
	ReportFieldWiFiFirmwareVersion  int16 = 0x34 // string
	ReportFieldWiFiMCUVersion       int16 = 0x35 // string
	ReportFieldCO2CalibrationActive int16 = 0x4A // uint8, 1 for active now, 0 for not active
	ReportFieldPMModuleSN           int16 = 0x61 // string, may be zero-length if no PM module present
	ReportFieldBatteryLevel         int16 = 0x64 // uint8, percentage
	ReportFieldSignalStrength       int16 = 0x65 // uint16
	// JSON
	ReportFieldMAC             string = "mac"        // string
	ReportFieldFirmwareVersion string = "sw_version" // string
	ReportFieldWiFiInfo        string = "wifi_info"  // string with format `SSID,RSSI,channel,BSSID`
	ReportFieldSensorData      string = "sensorData" // list of structures; note that spec says `sensor_data`
)
