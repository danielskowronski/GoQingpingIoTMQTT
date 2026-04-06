// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

const (
	// HEX
	VerbDataUploading          int16 = 0x31
	VerbConfigurationSending   int16 = 0x32
	VerbFirmwareUpgrade        int16 = 0x33
	VerbEventReporting         int16 = 0x34
	VerbConfigurationReporting int16 = 0x39
	VerbNetworkAccessSetting   int16 = 0x3A
	VerbRealtimeDataUploading  int16 = 0x3B
	// JSON
	VerbBleConnectionRequest          int16 = 1  // "Server sends BLE connection request"
	VerbBleDisconnectionRequest       int16 = 2  // "Server sends BLE disconnection request"
	VerbOpenBleNotificationRequest    int16 = 3  // "Server sends open BLE notification request"
	VerbCloseBleNotificationRequest   int16 = 4  // "Server sends close BLE notification request"
	VerbBleNotificationResponse       int16 = 5  // "Response for BLE notification from device"
	VerbBleDataWithResponse           int16 = 6  // "Server sends BLE data(with response)"
	VerbBleDataRead                   int16 = 7  // "Server reads BLE data"
	VerbBleDataResponse               int16 = 8  // "Response for BLE data from device"
	VerbBroadcastData                 int16 = 9  // "Broadcast data from device"
	VerbDeviceListRequest             int16 = 10 // "Device requests device list"
	VerbDeviceListResponse            int16 = 11 // "Server response device list"
	VerbTemporaryReportSetting        int16 = 12 // "Server send setting for temporary report and duration time"
	VerbHeartbeat                     int16 = 13 // "Heartbeat package"
	VerbReconnectMqtt                 int16 = 14 // "Reconnect MQTT"
	VerbBleDataWithoutResponse        int16 = 15 // "Server sends BLE data(without response)"
	VerbModifyMqttConnectionSetting   int16 = 16 // "Modify MQTT connection setting"
	VerbModifyDataReportInterval      int16 = 17 // "Modify data report interval"
	VerbSensorDataReportRealtime      int16 = 12 // "Sensor data report - real-time data"
	VerbSensorDataReportHistory       int16 = 17 // "Sensor data report - history data"
	VerbHistoryDataReportResponse     int16 = 18 // "Server response for history data report in type 17"
	VerbDeviceLogReport               int16 = 19 // "Device report log"
	VerbBindingStatus                 int16 = 20 // "Binding status"
	VerbOtaCommand                    int16 = 23 // "Server sends OTA command"
	VerbOtaCommandResponse            int16 = 24 // "Devices response for OTA command"
	VerbDeviceListRequestWithName     int16 = 25 // "Device requests device list(with device name)"
	VerbDeviceListResponseWithName    int16 = 26 // "Server responses device list(with device name)"
	VerbBindingStatusThirdPartyDevice int16 = 27 // "Binding status for third part's device"
	VerbReadDeviceSettingRequest      int16 = 28 // "Request to read device setting"
)
