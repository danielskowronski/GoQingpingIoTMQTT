// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

// CmdField... contain write-only fields

const (
	// HEX
	CmdFieldFirmwareUrlVersion         int16 = 0x11 // string
	CmdFieldFirmwareUrlWiFi            int16 = 0x12 // string, HTTP
	CmdFieldFirmwareUrlMCU             int16 = 0x13 // string, HTTP
	CmdFieldFactoryReset               int16 = 0x1E // ??
	CmdFieldWiFiParametersHEX          int16 = 0x20 // string, `"SSID","PASSWORD"`
	CmdFieldMQTTParametersHEX          int16 = 0x25 // string, `URL port username password clientid sub_topic pub_topic`
	CmdFieldNBIoTDataSplicing          int16 = 0x26 // 3 bytes (?), see spec
	CmdFieldNBIoTProtocolByteReplacing int16 = 0x27 // 3 bytes, see spec
	CmdFieldDataEncryption             int16 = 0x28 // 16 bytes, 0 for no encryption, otherwise AES-128 key
	CmdFieldBluetoothBroadcastName     int16 = 0x36 // string, max 28 bytes
	CmdFieldCO2CalibrateNow            int16 = 0x41 // uint16, write 1 to calibrate now
	CmdFieldRapidReporting             int16 = 0x42 // uint16, time in seconds
	CmdFieldSNTPServer                 int16 = 0x43 // string, SNTP server host
	CmdFieldSNTPEnabled                int16 = 0x44 // uint16, 0 for disabled, 1 for enabled

	// HEX - unsure if those are Cmd or Setting fields, may need to move them
	CmdFieldBluetoothBroadcastPower int16 = 0x2B // ??
	CmdFieldTime12Hours             int16 = 0x3E // uint16, 0 for 24 hours, 1 for 12 hours
	CmdFieldTVOCDisplayEnabled      int16 = 0x62 // uint8, 0 for disabled, 1 for enabled
	CmdFieldLEDEnabled              int16 = 0x63 // uint8, 0 for disabled, 1 for enabled

	// JSON
	CmdFieldRapidUploadInterval string = "up_itvl"  // uint16, seconds
	CmdFieldRapidUploadDuration string = "duration" // uint16, seconds
	CmdFieldMQTTParametersJSON  string = "mqtt_cfg" // object with fields: host, port, usrname, password, clientid, subscribe_topic, publish_topic
	CmdFieldWiFiParametersJSON  string = "wifi_cfg" // object with fields: ssid, password
)
