// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

type VerbInfo struct {
	Protocol  ProtocolType
	Direction MessageDirection
	Type      MessageType
}

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
	Verb12                            int16 = 12
	Verb17                            int16 = 17
	Verb28                            int16 = 28
	VerbBleConnectionRequest          int16 = 1      // "Server sends BLE connection request"
	VerbBleDisconnectionRequest       int16 = 2      // "Server sends BLE disconnection request"
	VerbOpenBleNotificationRequest    int16 = 3      // "Server sends open BLE notification request"
	VerbCloseBleNotificationRequest   int16 = 4      // "Server sends close BLE notification request"
	VerbBleNotificationResponse       int16 = 5      // "Response for BLE notification from device"
	VerbBleDataWithResponse           int16 = 6      // "Server sends BLE data(with response)"
	VerbBleDataRead                   int16 = 7      // "Server reads BLE data"
	VerbBleDataResponse               int16 = 8      // "Response for BLE data from device"
	VerbBroadcastData                 int16 = 9      // "Broadcast data from device"
	VerbDeviceListRequest             int16 = 10     // "Device requests device list"
	VerbDeviceListResponse            int16 = 11     // "Server response device list"
	VerbTemporaryReportSetting        int16 = Verb12 // "Server send setting for temporary report and duration time"
	VerbHeartbeat                     int16 = 13     // "Heartbeat package"
	VerbReconnectMqtt                 int16 = 14     // "Reconnect MQTT"
	VerbBleDataWithoutResponse        int16 = 15     // "Server sends BLE data(without response)"
	VerbModifyMqttConnectionSetting   int16 = 16     // "Modify MQTT connection setting"
	VerbModifyDataReportInterval      int16 = Verb17 // "Modify data report interval"
	VerbSensorDataReportRealtime      int16 = Verb12 // "Sensor data report - real-time data"
	VerbSensorDataReportHistory       int16 = Verb17 // "Sensor data report - history data"
	VerbHistoryDataReportResponse     int16 = 18     // "Server response for history data report in type 17"
	VerbDeviceLogReport               int16 = 19     // "Device report log"
	VerbBindingStatus                 int16 = 20     // "Binding status"
	VerbOtaCommand                    int16 = 23     // "Server sends OTA command"
	VerbOtaCommandResponse            int16 = 24     // "Devices response for OTA command"
	VerbDeviceListRequestWithName     int16 = 25     // "Device requests device list(with device name)"
	VerbDeviceListResponseWithName    int16 = 26     // "Server responses device list(with device name)"
	VerbBindingStatusThirdPartyDevice int16 = 27     // "Binding status for third part's device"
	VerbReadDeviceSettingRequest      int16 = Verb28 // "Request to read device setting"
)

var VerbInfoMap = map[int16]VerbInfo{
	// HEX
	VerbDataUploading:          {Protocol: ProtocolHEX, Direction: MessageDirectionDeviceToServer, Type: MessageTypeMeasurement},
	VerbConfigurationSending:   {Protocol: ProtocolHEX, Direction: MessageDirectionServerToDevice, Type: MessageTypeRequestFromServer},
	VerbFirmwareUpgrade:        {Protocol: ProtocolHEX, Direction: MessageDirectionServerToDevice, Type: MessageTypeRequestFromServer},
	VerbEventReporting:         {Protocol: ProtocolHEX, Direction: MessageDirectionDeviceToServer, Type: MessageTypeMeasurement},
	VerbConfigurationReporting: {Protocol: ProtocolHEX, Direction: MessageDirectionDeviceToServer, Type: MessageTypeResponseFromDevice},
	VerbNetworkAccessSetting:   {Protocol: ProtocolHEX, Direction: MessageDirectionServerToDevice, Type: MessageTypeRequestFromServer},
	VerbRealtimeDataUploading:  {Protocol: ProtocolHEX, Direction: MessageDirectionDeviceToServer, Type: MessageTypeMeasurement},
	// JSON
	VerbBleConnectionRequest:        {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbBleDisconnectionRequest:     {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbOpenBleNotificationRequest:  {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbCloseBleNotificationRequest: {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbBleNotificationResponse:     {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeOtherFromDevice},

	VerbBleDataWithResponse: {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbBleDataRead:         {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbBleDataResponse:     {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeOtherFromDevice},
	VerbBroadcastData:       {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeOtherFromDevice},
	VerbDeviceListRequest:   {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeOtherFromDevice},

	VerbDeviceListResponse:            {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbHeartbeat:                     {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeOtherFromDevice},
	VerbReconnectMqtt:                 {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeRequestFromServer},
	VerbBleDataWithoutResponse:        {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbModifyMqttConnectionSetting:   {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeRequestFromServer},
	VerbHistoryDataReportResponse:     {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeAck},
	VerbDeviceLogReport:               {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeOtherFromDevice},
	VerbBindingStatus:                 {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeOtherFromDevice},
	VerbOtaCommand:                    {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeRequestFromServer},
	VerbOtaCommandResponse:            {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeResponseFromDevice},
	VerbDeviceListRequestWithName:     {Protocol: ProtocolJSON, Direction: MessageDirectionDeviceToServer, Type: MessageTypeOtherFromDevice},
	VerbDeviceListResponseWithName:    {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
	VerbBindingStatusThirdPartyDevice: {Protocol: ProtocolJSON, Direction: MessageDirectionServerToDevice, Type: MessageTypeOtherFromServer},
}
